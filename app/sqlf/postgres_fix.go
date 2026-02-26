package sqlf

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"io"
	"strconv"
	"strings"

	"github.com/lib/pq"
)

func init() {
	sql.Register("postgres_fix", &postgresFixDriver{})
}

type postgresFixDriver struct {
	internal pq.Driver
}

func (d *postgresFixDriver) Open(name string) (driver.Conn, error) {
	conn, err := d.internal.Open(name)
	if err != nil {
		return nil, err
	}
	return &postgresFixConn{Conn: conn}, nil
}

type postgresFixConn struct {
	driver.Conn
}

func (c *postgresFixConn) Prepare(query string) (driver.Stmt, error) {
	rewritten := rewritePostgresSQL(query)
	isInsert := isInsertWithoutReturning(rewritten)
	if isInsert {
		rewritten += " RETURNING *"
	}
	stmt, err := c.Conn.Prepare(rewritten)
	if err != nil {
		return nil, err
	}
	return &postgresFixStmt{
		Stmt:     stmt,
		isInsert: isInsert,
	}, nil
}

func (c *postgresFixConn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	rewritten := rewritePostgresSQL(query)
	isInsert := isInsertWithoutReturning(rewritten)
	if isInsert {
		rewritten += " RETURNING *"
	}
	if prepareCtx, ok := c.Conn.(driver.ConnPrepareContext); ok {
		stmt, err := prepareCtx.PrepareContext(ctx, rewritten)
		if err != nil {
			return nil, err
		}
		return &postgresFixStmt{
			Stmt:     stmt,
			isInsert: isInsert,
		}, nil
	}
	return c.Prepare(query)
}

func (c *postgresFixConn) Exec(query string, args []driver.Value) (driver.Result, error) {
	rewritten := rewritePostgresSQL(query)
	if isInsertWithoutReturning(rewritten) {
		return c.execInsert(rewritten+" RETURNING *", args)
	}
	if execer, ok := c.Conn.(driver.Execer); ok {
		return execer.Exec(rewritten, args)
	}
	return nil, driver.ErrSkip
}

func (c *postgresFixConn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	rewritten := rewritePostgresSQL(query)
	if isInsertWithoutReturning(rewritten) {
		return c.execInsertContext(ctx, rewritten+" RETURNING *", args)
	}
	if execerCtx, ok := c.Conn.(driver.ExecerContext); ok {
		return execerCtx.ExecContext(ctx, rewritten, args)
	}
	return nil, driver.ErrSkip
}

func (c *postgresFixConn) Query(query string, args []driver.Value) (driver.Rows, error) {
	rewritten := rewritePostgresSQL(query)
	if queryer, ok := c.Conn.(driver.Queryer); ok {
		return queryer.Query(rewritten, args)
	}
	return nil, driver.ErrSkip
}

func (c *postgresFixConn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	rewritten := rewritePostgresSQL(query)
	if queryerCtx, ok := c.Conn.(driver.QueryerContext); ok {
		return queryerCtx.QueryContext(ctx, rewritten, args)
	}
	return nil, driver.ErrSkip
}

func (c *postgresFixConn) execInsert(query string, args []driver.Value) (driver.Result, error) {
	if queryer, ok := c.Conn.(driver.Queryer); ok {
		rows, err := queryer.Query(query, args)
		if err != nil {
			return nil, err
		}
		return parseInsertResult(rows)
	}
	return nil, driver.ErrSkip
}

func (c *postgresFixConn) execInsertContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	if queryerCtx, ok := c.Conn.(driver.QueryerContext); ok {
		rows, err := queryerCtx.QueryContext(ctx, query, args)
		if err != nil {
			return nil, err
		}
		return parseInsertResult(rows)
	}
	if queryer, ok := c.Conn.(driver.Queryer); ok {
		dargs, err := postgresNamedValueToValue(args)
		if err != nil {
			return nil, err
		}
		rows, err := queryer.Query(query, dargs)
		if err != nil {
			return nil, err
		}
		return parseInsertResult(rows)
	}
	return nil, driver.ErrSkip
}

type postgresFixStmt struct {
	driver.Stmt
	isInsert bool
}

func (s *postgresFixStmt) Exec(args []driver.Value) (driver.Result, error) {
	if s.isInsert == false {
		return s.Stmt.Exec(args)
	}
	rows, err := s.Stmt.Query(args)
	if err != nil {
		return nil, err
	}
	return parseInsertResult(rows)
}

func (s *postgresFixStmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	if s.isInsert == false {
		if stmtCtx, ok := s.Stmt.(driver.StmtExecContext); ok {
			return stmtCtx.ExecContext(ctx, args)
		}
		dargs, err := postgresNamedValueToValue(args)
		if err != nil {
			return nil, err
		}
		return s.Stmt.Exec(dargs)
	}
	if stmtCtx, ok := s.Stmt.(driver.StmtQueryContext); ok {
		rows, err := stmtCtx.QueryContext(ctx, args)
		if err != nil {
			return nil, err
		}
		return parseInsertResult(rows)
	}
	dargs, err := postgresNamedValueToValue(args)
	if err != nil {
		return nil, err
	}
	rows, err := s.Stmt.Query(dargs)
	if err != nil {
		return nil, err
	}
	return parseInsertResult(rows)
}

func (s *postgresFixStmt) Query(args []driver.Value) (driver.Rows, error) {
	return s.Stmt.Query(args)
}

func (s *postgresFixStmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	if stmtCtx, ok := s.Stmt.(driver.StmtQueryContext); ok {
		return stmtCtx.QueryContext(ctx, args)
	}
	dargs, err := postgresNamedValueToValue(args)
	if err != nil {
		return nil, err
	}
	return s.Stmt.Query(dargs)
}

type postgresFixResult struct {
	lastInsertID int64
	hasInsertID  bool
	rowsAffected int64
}

func (r *postgresFixResult) LastInsertId() (int64, error) {
	if r.hasInsertID == false {
		return 0, errors.New("last insert id is unavailable")
	}
	return r.lastInsertID, nil
}

func (r *postgresFixResult) RowsAffected() (int64, error) {
	return r.rowsAffected, nil
}

func parseInsertResult(rows driver.Rows) (driver.Result, error) {
	defer rows.Close()

	columns := rows.Columns()
	if len(columns) == 0 {
		return nil, errors.New("insert returning has no columns")
	}
	values := make([]driver.Value, len(columns))
	rowsAffected := int64(0)
	lastInsertID := int64(0)
	hasInsertID := false

	for {
		err := rows.Next(values)
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if hasInsertID == false {
			id, ok := toInt64(values[0])
			if ok {
				lastInsertID = id
				hasInsertID = true
			}
		}
		rowsAffected++
	}
	return &postgresFixResult{
		lastInsertID: lastInsertID,
		hasInsertID:  hasInsertID,
		rowsAffected: rowsAffected,
	}, nil
}

func toInt64(data interface{}) (int64, bool) {
	switch v := data.(type) {
	case int64:
		return v, true
	case int32:
		return int64(v), true
	case int:
		return int64(v), true
	case []byte:
		value, err := strconv.ParseInt(string(v), 10, 64)
		if err != nil {
			return 0, false
		}
		return value, true
	case string:
		value, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return 0, false
		}
		return value, true
	default:
		return 0, false
	}
}

func postgresNamedValueToValue(named []driver.NamedValue) ([]driver.Value, error) {
	dargs := make([]driver.Value, len(named))
	for i, nv := range named {
		if nv.Name != "" {
			return nil, errors.New("named parameters are not supported")
		}
		dargs[i] = nv.Value
	}
	return dargs, nil
}

func isInsertWithoutReturning(query string) bool {
	lowerQuery := strings.ToLower(strings.TrimSpace(query))
	return strings.HasPrefix(lowerQuery, "insert ") &&
		strings.Contains(lowerQuery, " returning ") == false
}

func rewritePostgresSQL(query string) string {
	builder := strings.Builder{}
	builder.Grow(len(query) + 16)
	placeholderIndex := 1
	inSingleQuote := false
	inDoubleQuote := false

	for i := 0; i < len(query); i++ {
		ch := query[i]

		if inSingleQuote {
			builder.WriteByte(ch)
			if ch == '\'' {
				if i+1 < len(query) && query[i+1] == '\'' {
					i++
					builder.WriteByte(query[i])
				} else {
					inSingleQuote = false
				}
			}
			continue
		}
		if inDoubleQuote {
			builder.WriteByte(ch)
			if ch == '"' {
				if i+1 < len(query) && query[i+1] == '"' {
					i++
					builder.WriteByte(query[i])
				} else {
					inDoubleQuote = false
				}
			}
			continue
		}

		switch ch {
		case '\'':
			inSingleQuote = true
			builder.WriteByte(ch)
		case '"':
			inDoubleQuote = true
			builder.WriteByte(ch)
		case '`':
			builder.WriteByte('"')
		case '?':
			builder.WriteByte('$')
			builder.WriteString(strconv.Itoa(placeholderIndex))
			placeholderIndex++
		default:
			builder.WriteByte(ch)
		}
	}
	return builder.String()
}
