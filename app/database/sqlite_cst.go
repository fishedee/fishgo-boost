package database

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"time"

	"modernc.org/sqlite" // 引用原版驱动
)

// 定义东八区时区（避免依赖系统的 tzdata，兼容 Windows/Docker）
var eightHourSeconds = 8 * 3600

var cstZone = time.FixedZone("CST", eightHourSeconds)

func init() {
	// 注册一个新的驱动名：sqlite_cst
	sql.Register("sqlite_cst", &CstDriver{})
}

// CstDriver 包装了原生的 sqlite 驱动
type CstDriver struct {
	// 内部持有原版驱动实例
	internal sqlite.Driver
}

// Open 实现 driver.Driver 接口
func (d *CstDriver) Open(name string) (driver.Conn, error) {
	// 1. 【处理读取】自动在 DSN 后追加 _loc=Asia/Shanghai
	// 这样查询出来的 time.Time 会自动附带时区信息
	/*
		if !strings.Contains(name, "_loc=") {
			if strings.Contains(name, "?") {
				name += "&_loc=Asia/Shanghai"
			} else {
				name += "?_loc=Asia/Shanghai"
			}
		}
	*/
	// 2. 调用原版驱动打开连接
	conn, err := d.internal.Open(name)
	if err != nil {
		return nil, err
	}
	// 3. 【处理写入】返回一个包装过的 Connection
	return &cstConn{Conn: conn}, nil
}

// ---------------------------------------------------------
// 1. 连接包装 (处理写入 + 创建查询)
// ---------------------------------------------------------
type cstConn struct {
	driver.Conn
}

// 拦截 Prepare：这是 XORM 写入的关键路径
func (c *cstConn) Prepare(query string) (driver.Stmt, error) {
	stmt, err := c.Conn.Prepare(query)
	if err != nil {
		return nil, err
	}
	// 返回我们要包装过的 Stmt
	return &cstStmt{Stmt: stmt}, nil
}

// 拦截 PrepareContext (Go 1.8+)
func (c *cstConn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	if prepareCtx, ok := c.Conn.(driver.ConnPrepareContext); ok {
		stmt, err := prepareCtx.PrepareContext(ctx, query)
		if err != nil {
			return nil, err
		}
		return &cstStmt{Stmt: stmt}, nil
	}
	return c.Prepare(query)
}

// 拦截直接 Exec (非 Prepare 模式)
func (c *cstConn) Exec(query string, args []driver.Value) (driver.Result, error) {
	fixArgs(args) // 修改参数
	if execer, ok := c.Conn.(driver.Execer); ok {
		return execer.Exec(query, args)
	}
	return nil, driver.ErrSkip
}

// 拦截直接 ExecContext
func (c *cstConn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	fixNamedArgs(args) // 修改参数
	if execerCtx, ok := c.Conn.(driver.ExecerContext); ok {
		return execerCtx.ExecContext(ctx, query, args)
	}
	return nil, driver.ErrSkip
}

// 拦截 Query，返回我们自定义的 Rows
func (c *cstConn) Query(query string, args []driver.Value) (driver.Rows, error) {
	fixArgs(args)
	if queryer, ok := c.Conn.(driver.Queryer); ok {
		rows, err := queryer.Query(query, args)
		if err != nil {
			return nil, err
		}
		// 返回包装过的 Rows
		return &cstRows{Rows: rows}, nil
	}
	return nil, driver.ErrSkip
}

// 支持 Context 的 Query
func (c *cstConn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	fixNamedArgs(args)
	if queryerCtx, ok := c.Conn.(driver.QueryerContext); ok {
		rows, err := queryerCtx.QueryContext(ctx, query, args)
		if err != nil {
			return nil, err
		}
		// 返回包装过的 Rows
		return &cstRows{Rows: rows}, nil
	}
	return nil, driver.ErrSkip
}

// ---------------------------------------------------------
// 2. Statement 包装 (这是你之前缺失的部分)
// ---------------------------------------------------------
type cstStmt struct {
	driver.Stmt
}

// 拦截 Stmt 的 Exec：XORM 绝大多数写入走这里
func (s *cstStmt) Exec(args []driver.Value) (driver.Result, error) {
	fixArgs(args) // <--- 关键：这里拦截写入参数
	return s.Stmt.Exec(args)
}

// 拦截 Stmt 的 ExecContext
func (s *cstStmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	fixNamedArgs(args) // <--- 关键：这里拦截写入参数
	if stmtCtx, ok := s.Stmt.(driver.StmtExecContext); ok {
		return stmtCtx.ExecContext(ctx, args)
	}
	// Fallback logic if needed, though modernc supports ExecContext
	dargs, _ := namedValueToValue(args)
	return s.Stmt.Exec(dargs)
}

// 拦截 Stmt 的 Query (读取)
func (s *cstStmt) Query(args []driver.Value) (driver.Rows, error) {
	fixArgs(args)
	rows, err := s.Stmt.Query(args)
	if err != nil {
		return nil, err
	}
	return &cstRows{Rows: rows}, nil
}

// 拦截 Stmt 的 QueryContext
func (s *cstStmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	fixNamedArgs(args)
	if stmtCtx, ok := s.Stmt.(driver.StmtQueryContext); ok {
		rows, err := stmtCtx.QueryContext(ctx, args)
		if err != nil {
			return nil, err
		}
		return &cstRows{Rows: rows}, nil
	}
	dargs, _ := namedValueToValue(args)
	rows, err := s.Stmt.Query(dargs)
	if err != nil {
		return nil, err
	}
	return &cstRows{Rows: rows}, nil
}

// Close 关闭语句
func (s *cstStmt) Close() error {
	return s.Stmt.Close()
}

// NumInput 返回占位符数量
func (s *cstStmt) NumInput() int {
	return s.Stmt.NumInput()
}

// ---------------------------------------------------------
// 3. Rows 包装 (处理读取) -- 核心部分
// ---------------------------------------------------------
type cstRows struct {
	driver.Rows
}

// Next 是读取数据的核心方法
func (r *cstRows) Next(dest []driver.Value) error {
	// 1. 先让底层驱动把数据读出来放到 dest 里
	err := r.Rows.Next(dest)
	if err != nil {
		return err
	}
	// 2. 遍历结果，如果是 time.Time，强制转回 CST
	for i, v := range dest {
		if t, ok := v.(time.Time); ok {
			// 这里的 In(cstZone) 做了两件事：
			// 1. 如果它是 UTC，会加上 8小时显示。
			// 2. 将 Location 信息强制改为 CST。
			dest[i] = time.Unix((t.Unix() - int64(eightHourSeconds)), 0).In(cstZone)
		}
	}
	return nil
}

// Columns 返回列名
func (r *cstRows) Columns() []string {
	return r.Rows.Columns()
}

// Close 关闭结果集
func (r *cstRows) Close() error {
	return r.Rows.Close()
}

func fixArgs(args []driver.Value) {
	//传入的datetime是string类型，所以这一段逻辑忽略
	/*
		for i, v := range args {
			if t, ok := v.(time.Time); ok {
				// 核心逻辑：无论传入什么时区，都转为 CST
				args[i] = t.UTC()
			}
		}
	*/
}

// 修正命名参数 (Context 方法用)
func fixNamedArgs(args []driver.NamedValue) {
	//传入的datetime是string类型，所以这一段逻辑忽略
	/*
		for i := range args {
			if t, ok := args[i].Value.(time.Time); ok {
				args[i].Value = t.UTC()
			}
		}
	*/
}

func namedValueToValue(named []driver.NamedValue) ([]driver.Value, error) {
	dargs := make([]driver.Value, len(named))
	for i, nv := range named {
		dargs[i] = nv.Value
	}
	return dargs, nil
}
