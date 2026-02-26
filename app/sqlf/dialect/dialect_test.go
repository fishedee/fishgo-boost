package dialect

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
	AppendQuotedIdentifier(GetSqlDialect("mysql"), "userId", &builder)
	if builder.String() != "`userId`" {
		t.Fatalf("mysql quote mismatch: %v", builder.String())
	}

	builder.Reset()
	AppendQuotedIdentifier(GetSqlDialect("sqlite_fix"), "userId", &builder)
	if builder.String() != "`userId`" {
		t.Fatalf("sqlite quote mismatch: %v", builder.String())
	}

	builder.Reset()
	AppendQuotedIdentifier(GetSqlDialect("postgres_fix"), "userId", &builder)
	if builder.String() != "\"userid\"" {
		t.Fatalf("postgres quote mismatch: %v", builder.String())
	}
}

func TestDialectLimit(t *testing.T) {
	if GetSqlDialect("mysql").Limit(0, 10) != "limit 0,10" {
		t.Fatalf("mysql limit mismatch")
	}
	if GetSqlDialect("sqlite_fix").Limit(0, 10) != "limit 0,10" {
		t.Fatalf("sqlite limit mismatch")
	}
	if GetSqlDialect("postgres_fix").Limit(0, 10) != "limit 10 offset 0" {
		t.Fatalf("postgres limit mismatch")
	}
}
