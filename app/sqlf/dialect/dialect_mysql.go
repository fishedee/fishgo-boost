package dialect

import (
	"fmt"
	"reflect"
	"strings"
)

type mysqlDialect struct{}

func (d mysqlDialect) Engine() ENGINE {
	return MYSQL
}

func (d mysqlDialect) QuoteIdentifier(name string, builder *strings.Builder) {
	builder.WriteByte('`')
	builder.WriteString(name)
	builder.WriteByte('`')
}

func (d mysqlDialect) Limit(index int, size int) string {
	return fmt.Sprintf("limit %d,%d", index, size)
}

func (d mysqlDialect) NeedFromResultConvert(inValueType reflect.Type) bool {
	return false
}

func (d mysqlDialect) FromResultConvertValue(inValue reflect.Value) (reflect.Value, fromResultConvertValueFunc) {
	err := fmt.Sprintf("mysql do not need to fromResultConvertValue:%s", inValue.Type())
	panic(err)
}

func (d mysqlDialect) NeedToArgsConvert(inValueType reflect.Type) bool {
	return false
}

func (d mysqlDialect) ToArgsConvertValue(inValue reflect.Value) interface{} {
	err := fmt.Sprintf("mysql do not need to toArgsConvertValue:%s", inValue.Type())
	panic(err)
}
