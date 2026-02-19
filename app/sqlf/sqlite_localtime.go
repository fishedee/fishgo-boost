package sqlf

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"time"

	"modernc.org/sqlite"
)

func init() {
	sql.Register("sqlite_localtime", &sqliteLocaltimeDriver{})
}

type sqliteLocaltimeDriver struct {
	internal sqlite.Driver
}

func (d *sqliteLocaltimeDriver) Open(name string) (driver.Conn, error) {
	conn, err := d.internal.Open(name)
	if err != nil {
		return nil, err
	}
	return &sqliteLocaltimeConn{Conn: conn}, nil
}

type sqliteLocaltimeConn struct {
	driver.Conn
}

func (c *sqliteLocaltimeConn) Prepare(query string) (driver.Stmt, error) {
	stmt, err := c.Conn.Prepare(query)
	if err != nil {
		return nil, err
	}
	return &sqliteLocaltimeStmt{Stmt: stmt}, nil
}

func (c *sqliteLocaltimeConn) PrepareContext(ctx context.Context, query string) (driver.Stmt, error) {
	if prepareCtx, ok := c.Conn.(driver.ConnPrepareContext); ok {
		stmt, err := prepareCtx.PrepareContext(ctx, query)
		if err != nil {
			return nil, err
		}
		return &sqliteLocaltimeStmt{Stmt: stmt}, nil
	}
	return c.Prepare(query)
}

func (c *sqliteLocaltimeConn) Exec(query string, args []driver.Value) (driver.Result, error) {
	convertArgsToUTC(args)
	if execer, ok := c.Conn.(driver.Execer); ok {
		return execer.Exec(query, args)
	}
	return nil, driver.ErrSkip
}

func (c *sqliteLocaltimeConn) ExecContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Result, error) {
	convertNamedArgsToUTC(args)
	if execerCtx, ok := c.Conn.(driver.ExecerContext); ok {
		return execerCtx.ExecContext(ctx, query, args)
	}
	return nil, driver.ErrSkip
}

func (c *sqliteLocaltimeConn) Query(query string, args []driver.Value) (driver.Rows, error) {
	convertArgsToUTC(args)
	if queryer, ok := c.Conn.(driver.Queryer); ok {
		rows, err := queryer.Query(query, args)
		if err != nil {
			return nil, err
		}
		return &sqliteLocaltimeRows{Rows: rows}, nil
	}
	return nil, driver.ErrSkip
}

func (c *sqliteLocaltimeConn) QueryContext(ctx context.Context, query string, args []driver.NamedValue) (driver.Rows, error) {
	convertNamedArgsToUTC(args)
	if queryerCtx, ok := c.Conn.(driver.QueryerContext); ok {
		rows, err := queryerCtx.QueryContext(ctx, query, args)
		if err != nil {
			return nil, err
		}
		return &sqliteLocaltimeRows{Rows: rows}, nil
	}
	return nil, driver.ErrSkip
}

type sqliteLocaltimeStmt struct {
	driver.Stmt
}

func (s *sqliteLocaltimeStmt) Exec(args []driver.Value) (driver.Result, error) {
	convertArgsToUTC(args)
	return s.Stmt.Exec(args)
}

func (s *sqliteLocaltimeStmt) ExecContext(ctx context.Context, args []driver.NamedValue) (driver.Result, error) {
	convertNamedArgsToUTC(args)
	if stmtCtx, ok := s.Stmt.(driver.StmtExecContext); ok {
		return stmtCtx.ExecContext(ctx, args)
	}
	dargs, _ := namedValueToValue(args)
	return s.Stmt.Exec(dargs)
}

func (s *sqliteLocaltimeStmt) Query(args []driver.Value) (driver.Rows, error) {
	convertArgsToUTC(args)
	rows, err := s.Stmt.Query(args)
	if err != nil {
		return nil, err
	}
	return &sqliteLocaltimeRows{Rows: rows}, nil
}

func (s *sqliteLocaltimeStmt) QueryContext(ctx context.Context, args []driver.NamedValue) (driver.Rows, error) {
	convertNamedArgsToUTC(args)
	if stmtCtx, ok := s.Stmt.(driver.StmtQueryContext); ok {
		rows, err := stmtCtx.QueryContext(ctx, args)
		if err != nil {
			return nil, err
		}
		return &sqliteLocaltimeRows{Rows: rows}, nil
	}
	dargs, _ := namedValueToValue(args)
	rows, err := s.Stmt.Query(dargs)
	if err != nil {
		return nil, err
	}
	return &sqliteLocaltimeRows{Rows: rows}, nil
}

type sqliteLocaltimeRows struct {
	driver.Rows
}

func (r *sqliteLocaltimeRows) Next(dest []driver.Value) error {
	if err := r.Rows.Next(dest); err != nil {
		return err
	}
	convertDestToLocal(dest)
	return nil
}

func convertArgsToUTC(args []driver.Value) {
	for i, v := range args {
		if t, ok := v.(time.Time); ok {
			args[i] = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)
		}
	}
}

func convertNamedArgsToUTC(args []driver.NamedValue) {
	for i := range args {
		if t, ok := args[i].Value.(time.Time); ok {
			args[i].Value = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.UTC)
		}
	}
}

func convertDestToLocal(dest []driver.Value) {
	for i, v := range dest {
		if t, ok := v.(time.Time); ok {
			dest[i] = time.Date(t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), time.Local)
		}
	}
}

func namedValueToValue(named []driver.NamedValue) ([]driver.Value, error) {
	dargs := make([]driver.Value, len(named))
	for i, nv := range named {
		dargs[i] = nv.Value
	}
	return dargs, nil
}
