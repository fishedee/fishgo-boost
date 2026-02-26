package dialect

import (
	"encoding/json"
	"reflect"
	"strings"
	"time"
)

type fromResultConvertValueFunc = func()

type ENGINE int

const (
	MYSQL ENGINE = 1 + iota
	SQLITE
	POSTGRES
	UNKNOWN
)

type SqlDialect interface {
	Engine() ENGINE
	QuoteIdentifier(name string, builder *strings.Builder)
	Limit(index int, size int) string
	NeedFromResultConvert(inValueType reflect.Type) bool
	FromResultConvertValue(inValue reflect.Value) (reflect.Value, fromResultConvertValueFunc)
	NeedToArgsConvert(inValueType reflect.Type) bool
	ToArgsConvertValue(inValue reflect.Value) interface{}
}

var (
	jsonRawMessageType = reflect.TypeOf(json.RawMessage{})
	timeType           = reflect.TypeOf(time.Time{})
)

var (
	mysqlSQLDialect    = mysqlDialect{}
	sqliteSQLDialect   = sqliteDialect{}
	postgresSQLDialect = postgresDialect{}
	unknownSQLDialect  = unknownDialect{}
)

func GetSqlDialect(driver string) SqlDialect {
	if strings.Contains(driver, "mysql") {
		return mysqlSQLDialect
	}
	if strings.Contains(driver, "sqlite") {
		return sqliteSQLDialect
	}
	if strings.Contains(driver, "postgres") {
		return postgresSQLDialect
	}
	return unknownSQLDialect
}

func AppendQuotedIdentifier(dialect SqlDialect, name string, builder *strings.Builder) {
	dialect.QuoteIdentifier(name, builder)
}
