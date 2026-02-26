package sqlf

import (
	"strings"
	"testing"
)

type dialectUser struct {
	UserId int `sqlf:"autoincr"`
	Name   string
}

func TestDialectQuoteIdentifier(t *testing.T) {
	builder := strings.Builder{}
	appendQuotedIdentifier(getSqlDialect("mysql"), "userId", &builder)
	if builder.String() != "`userId`" {
		t.Fatalf("mysql quote mismatch: %v", builder.String())
	}

	builder.Reset()
	appendQuotedIdentifier(getSqlDialect("sqlite_fix"), "userId", &builder)
	if builder.String() != "`userId`" {
		t.Fatalf("sqlite quote mismatch: %v", builder.String())
	}

	builder.Reset()
	appendQuotedIdentifier(getSqlDialect("postgres_fix"), "userId", &builder)
	if builder.String() != "\"userid\"" {
		t.Fatalf("postgres quote mismatch: %v", builder.String())
	}
}

func TestDialectLimit(t *testing.T) {
	if getSqlDialect("mysql").limit(0, 10) != "limit 0,10" {
		t.Fatalf("mysql limit mismatch")
	}
	if getSqlDialect("sqlite_fix").limit(0, 10) != "limit 0,10" {
		t.Fatalf("sqlite limit mismatch")
	}
	if getSqlDialect("postgres_fix").limit(0, 10) != "limit 10 offset 0" {
		t.Fatalf("postgres limit mismatch")
	}
}

func TestGenSqlUseDialectQuote(t *testing.T) {
	sql, args, err := genSql(getSqlDialect("postgres_fix"), "insert into t_user(?.insertColumn) values ?.insertValue", []interface{}{
		dialectUser{Name: "fish"},
		dialectUser{Name: "fish"},
	})
	if err != nil {
		t.Fatalf("genSql fail: %v", err)
	}
	if sql != "insert into t_user(\"name\") values (?)" {
		t.Fatalf("sql mismatch: %v", sql)
	}
	if len(args) != 1 || args[0] != "fish" {
		t.Fatalf("args mismatch: %v", args)
	}
}
