package typ

import (
	gosql "database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"

	. "github.com/fishedee/fishgo-boost/app/sqlf/dialect"
)

func dynamicType_fromResult(dialect SqlDialect, v interface{}, rows *gosql.Rows) error {
	inValue := reflect.ValueOf(v)
	inValueType := inValue.Type()
	if inValueType.Kind() != reflect.Ptr {
		return errors.New(fmt.Sprintf("%s dos not support setValue", inValueType))
	}
	inValueElem := inValue.Elem()
	inValueElemType := inValueElem.Type()
	if inValueElemType.Name() != "" || inValueElemType.Kind() != reflect.Slice {
		//单个转换
		if rows.Next() {
			wrapValue, elemValueWriter := dialect.FromResultConvertValue(inValueElem)
			err := rows.Scan(wrapValue.Addr().Interface())
			if err != nil {
				return err
			}
			elemValueWriter()
			return nil
		} else {
			return errors.New("has no result")
		}
	} else {
		//数组转换
		inValueElemTypeReal := inValueElemType.Elem()
		realResult := reflect.MakeSlice(reflect.SliceOf(inValueElemTypeReal), 0, 0)
		tempValue := reflect.New(inValueElemTypeReal).Elem()
		wrapValue, tempValueWriter := dialect.FromResultConvertValue(tempValue)
		for rows.Next() {
			err := rows.Scan(wrapValue.Addr().Interface())
			if err != nil {
				return err
			}
			tempValueWriter()
			realResult = reflect.Append(realResult, tempValue)
		}
		inValueElem.Set(realResult)
		return nil
	}
}

func dynamicType_toArgsfunc(dialect SqlDialect, isInsert bool, v interface{}, in []interface{}, builder *strings.Builder) ([]interface{}, error) {
	inValue := reflect.ValueOf(v)
	inValueType := inValue.Type()
	if inValueType.Kind() == reflect.Ptr {
		inValue = inValue.Elem()
		inValueType = inValueType.Elem()
	}
	if inValueType.Name() != "" || inValueType.Kind() != reflect.Slice {
		//单个转换
		tempValue := dialect.ToArgsConvertValue(inValue)
		builder.WriteByte('?')
		in = append(in, tempValue)
		return in, nil
	} else {
		//数组转换
		len := inValue.Len()
		builder.WriteString(getSqlComma(len))
		for i := 0; i != len; i++ {
			tempValue := dialect.ToArgsConvertValue(inValue.Index(i))
			in = append(in, tempValue)
		}
		return in, nil
	}
}
