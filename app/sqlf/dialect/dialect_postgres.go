package dialect

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

type postgresDialect struct{}

func (d postgresDialect) Engine() ENGINE {
	return POSTGRES
}

func (d postgresDialect) QuoteIdentifier(name string, builder *strings.Builder) {
	// Postgres folds unquoted identifiers to lower-case.
	builder.WriteByte('"')
	builder.WriteString(strings.ToLower(name))
	builder.WriteByte('"')
}

func (d postgresDialect) Limit(index int, size int) string {
	return fmt.Sprintf("limit %d offset %d", size, index)
}

func (d postgresDialect) NeedFromResultConvert(inValueType reflect.Type) bool {
	if inValueType == jsonRawMessageType {
		return true
	}
	return false
}

func (d postgresDialect) FromResultConvertValue(inValue reflect.Value) (reflect.Value, fromResultConvertValueFunc) {
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

func (d postgresDialect) NeedToArgsConvert(inValueType reflect.Type) bool {
	return false
}

func (d postgresDialect) ToArgsConvertValue(inValue reflect.Value) interface{} {
	err := fmt.Sprintf("postgresql do not need to toArgsConvertValue:%s", inValue.Type())
	panic(err)
}
