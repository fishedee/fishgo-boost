package sqlf

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type fromResultConvertValueFunc = func()

type sqlDialect interface {
	engine() ENGINE
	quoteIdentifier(name string, builder *strings.Builder)
	limit(index int, size int) string
	needFromResultConvert(inValueType reflect.Type) bool
	fromResultConvertValue(inValue reflect.Value) (reflect.Value, fromResultConvertValueFunc)
	needToArgsConvert(inValueType reflect.Type) bool
	toArgsConvertValue(inValue reflect.Value) interface{}
}

type mysqlDialect struct{}

func (d mysqlDialect) engine() ENGINE {
	return MYSQL
}

func (d mysqlDialect) quoteIdentifier(name string, builder *strings.Builder) {
	builder.WriteByte('`')
	builder.WriteString(name)
	builder.WriteByte('`')
}

func (d mysqlDialect) limit(index int, size int) string {
	return fmt.Sprintf("limit %d,%d", index, size)
}

func (d mysqlDialect) needFromResultConvert(inValueType reflect.Type) bool {
	return false
}

func (d mysqlDialect) fromResultConvertValue(inValue reflect.Value) (reflect.Value, fromResultConvertValueFunc) {
	err := fmt.Sprintf("mysql do not need to fromResultConvertValue:%s", inValue.Type())
	panic(err)
}

func (d mysqlDialect) needToArgsConvert(inValueType reflect.Type) bool {
	return false
}

func (d mysqlDialect) toArgsConvertValue(inValue reflect.Value) interface{} {
	err := fmt.Sprintf("mysql do not need to toArgsConvertValue:%s", inValue.Type())
	panic(err)
}

type sqliteDialect struct{}

func (d sqliteDialect) engine() ENGINE {
	return SQLITE
}

func (d sqliteDialect) quoteIdentifier(name string, builder *strings.Builder) {
	builder.WriteByte('`')
	builder.WriteString(name)
	builder.WriteByte('`')
}

func (d sqliteDialect) limit(index int, size int) string {
	return fmt.Sprintf("limit %d,%d", index, size)
}

func (d sqliteDialect) needFromResultConvert(inValueType reflect.Type) bool {
	return false
}

func (d sqliteDialect) fromResultConvertValue(inValue reflect.Value) (reflect.Value, fromResultConvertValueFunc) {
	err := fmt.Sprintf("sqlite do not need to fromResultConvertValue:%s", inValue.Type())
	panic(err)
}

func (d sqliteDialect) needToArgsConvert(inValueType reflect.Type) bool {
	return false
}

func (d sqliteDialect) toArgsConvertValue(inValue reflect.Value) interface{} {
	err := fmt.Sprintf("sqlite do not need to toArgsConvertValue:%s", inValue.Type())
	panic(err)
}

type postgresDialect struct{}

func (d postgresDialect) engine() ENGINE {
	return POSTGRES
}

func (d postgresDialect) quoteIdentifier(name string, builder *strings.Builder) {
	// Postgres folds unquoted identifiers to lower-case.
	builder.WriteByte('"')
	builder.WriteString(strings.ToLower(name))
	builder.WriteByte('"')
}

func (d postgresDialect) limit(index int, size int) string {
	return fmt.Sprintf("limit %d offset %d", size, index)
}

var (
	jsonRawMessageType = reflect.TypeOf(json.RawMessage{})
)

func (d postgresDialect) needFromResultConvert(inValueType reflect.Type) bool {
	if inValueType == jsonRawMessageType {
		return true
	}
	return false
}

func (d postgresDialect) fromResultConvertValue(inValue reflect.Value) (reflect.Value, fromResultConvertValueFunc) {
	if inValue.Type() == jsonRawMessageType {
		//读取json.RawMessage的时候，需要先读取[]byte，再转换为json.RawMessage来写入
		result := []byte{}
		resultValue := reflect.ValueOf(&result).Elem()
		resultValueWritter := func() {
			inValue.Set(reflect.ValueOf(json.RawMessage(result)))
		}
		return resultValue, resultValueWritter
	}
	err := fmt.Sprintf("postgresql do not need to fromResultConvertValue:%s", inValue.Type())
	panic(err)
}

func (d postgresDialect) needToArgsConvert(inValueType reflect.Type) bool {
	return false
}

func (d postgresDialect) toArgsConvertValue(inValue reflect.Value) interface{} {
	err := fmt.Sprintf("postgresql do not need to toArgsConvertValue:%s", inValue.Type())
	panic(err)
}

type unknownDialect struct{}

func (d unknownDialect) engine() ENGINE {
	return UNKNOWN
}

func (d unknownDialect) quoteIdentifier(name string, builder *strings.Builder) {
	builder.WriteByte('`')
	builder.WriteString(name)
	builder.WriteByte('`')
}

func (d unknownDialect) limit(index int, size int) string {
	return fmt.Sprintf("limit %d,%d", index, size)
}

func (d unknownDialect) needFromResultConvert(inValueType reflect.Type) bool {
	return false
}

func (d unknownDialect) fromResultConvertValue(inValue reflect.Value) (reflect.Value, fromResultConvertValueFunc) {
	err := fmt.Sprintf("unknownDialect do not need to fromResultConvertValue:%s", inValue.Type())
	panic(err)
}

func (d unknownDialect) needToArgsConvert(inValueType reflect.Type) bool {
	return false
}

func (d unknownDialect) toArgsConvertValue(inValue reflect.Value) interface{} {
	err := fmt.Sprintf("unknownDialect do not need to toArgsConvertValue:%s", inValue.Type())
	panic(err)
}

var (
	mysqlSQLDialect    = mysqlDialect{}
	sqliteSQLDialect   = sqliteDialect{}
	postgresSQLDialect = postgresDialect{}
	unknownSQLDialect  = unknownDialect{}
)

func getSqlDialect(driver string) sqlDialect {
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

func appendQuotedIdentifier(dialect sqlDialect, name string, builder *strings.Builder) {
	dialect.quoteIdentifier(name, builder)
}
