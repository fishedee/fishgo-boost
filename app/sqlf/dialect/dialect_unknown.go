package dialect

import (
	"fmt"
	"reflect"
	"strings"
)

type unknownDialect struct{}

func (d unknownDialect) Engine() ENGINE {
	return UNKNOWN
}

func (d unknownDialect) QuoteIdentifier(name string, builder *strings.Builder) {
	builder.WriteByte('`')
	builder.WriteString(name)
	builder.WriteByte('`')
}

func (d unknownDialect) Limit(index int, size int) string {
	return fmt.Sprintf("limit %d,%d", index, size)
}

func (d unknownDialect) NeedFromResultConvert(inValueType reflect.Type) bool {
	return false
}

func (d unknownDialect) FromResultConvertValue(inValue reflect.Value) (reflect.Value, fromResultConvertValueFunc) {
	err := fmt.Sprintf("unknownDialect do not need to fromResultConvertValue:%s", inValue.Type())
	panic(err)
}

func (d unknownDialect) NeedToArgsConvert(inValueType reflect.Type) bool {
	return false
}

func (d unknownDialect) ToArgsConvertValue(inValue reflect.Value) interface{} {
	err := fmt.Sprintf("unknownDialect do not need to toArgsConvertValue:%s", inValue.Type())
	panic(err)
}
