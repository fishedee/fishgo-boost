package dialect

import (
	"fmt"
	"reflect"
	"strings"
	"time"
)

type sqliteDialect struct{}

func (d sqliteDialect) Engine() ENGINE {
	return SQLITE
}

func (d sqliteDialect) QuoteIdentifier(name string, builder *strings.Builder) {
	builder.WriteByte('`')
	builder.WriteString(name)
	builder.WriteByte('`')
}

func (d sqliteDialect) Limit(index int, size int) string {
	return fmt.Sprintf("limit %d,%d", index, size)
}

func (d sqliteDialect) NeedFromResultConvert(inValueType reflect.Type) bool {
	return false
}

func (d sqliteDialect) FromResultConvertValue(inValue reflect.Value) (reflect.Value, fromResultConvertValueFunc) {
	err := fmt.Sprintf("sqlite do not need to fromResultConvertValue:%s", inValue.Type())
	panic(err)
}

func (d sqliteDialect) NeedToArgsConvert(inValueType reflect.Type) bool {
	if inValueType == timeType {
		return true
	}
	return false
}

func (d sqliteDialect) ToArgsConvertValue(inValue reflect.Value) interface{} {
	if inValue.Type() == timeType {
		t := inValue.Interface().(time.Time)
		return t.UTC().Format("2006-01-02 15:04:05")
	}
	err := fmt.Sprintf("sqlite do not need to toArgsConvertValue:%s", inValue.Type())
	panic(err)
}
