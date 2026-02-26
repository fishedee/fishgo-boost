package sqlf

import (
	gosql "database/sql"
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	. "github.com/fishedee/fishgo-boost/language"
)

func initStructTypeOperation(t reflect.Type) sqlTypeOperation {
	return sqlTypeOperation{
		toArgs:     initSqlToArgs(t),
		fromResult: initSqlFromResult(t),
		column:     initSqlColumn(t),
		setValue:   initSqlSetValue(t),
	}
}

func getTypeKind(t reflect.Type) int {
	timeType := reflect.TypeOf(time.Time{})
	if t.Kind() == reflect.Struct && t != timeType {
		//struct类型
		return 1
	} else if t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Struct && t.Elem() != timeType {
		//*struct类型
		return 2
	} else if t.Kind() == reflect.Slice && t.Elem().Kind() == reflect.Struct && t.Elem() != timeType {
		//[]struct类型
		return 3
	} else if t.Kind() == reflect.Ptr && t.Elem().Kind() == reflect.Slice && t.Elem().Elem().Kind() == reflect.Struct && t.Elem().Elem() != timeType {
		//*[]struct类型
		return 4
	} else {
		return 5
	}
}

type sqlStructPublicField struct {
	name       string
	fieldType  reflect.Type
	index      []int
	isTimeType bool
	isAutoIncr bool
	isCreated  bool
	isUpdated  bool
}

func getFieldInfo(field reflect.StructField) sqlStructPublicField {
	timeType := reflect.TypeOf(time.Time{})
	fieldName := field.Name
	fieldName = strings.ToLower(fieldName[0:1]) + fieldName[1:]
	fieldTag, isExist := field.Tag.Lookup("sqlf")
	if isExist == false {
		fieldTag = ""
	}
	tagList := Explode(fieldTag, ",")
	isAutoIncr := false
	isCreated := false
	isUpdated := false
	setNumber := 0
	for _, tag := range tagList {
		if tag == "autoincr" {
			isAutoIncr = true
			setNumber++
		}
		if tag == "created" {
			isCreated = true
			setNumber++
		}
		if tag == "updated" {
			isUpdated = true
			setNumber++
		}
	}
	if setNumber >= 2 {
		panic(fmt.Sprintf("only one tag specify %v.%v", field.PkgPath, field.Name))
	}
	nameTag, isExist := field.Tag.Lookup("sqlname")
	if isExist {
		fieldName = nameTag
	}
	return sqlStructPublicField{
		name:       fieldName,
		fieldType:  field.Type,
		index:      field.Index,
		isTimeType: field.Type == timeType,
		isAutoIncr: isAutoIncr,
		isCreated:  isCreated,
		isUpdated:  isUpdated,
	}
}
func getStructPublicField(t reflect.Type) []sqlStructPublicField {
	result := []sqlStructPublicField{}
	numField := t.NumField()
	for i := 0; i != numField; i++ {
		field := t.Field(i)
		fieldName := field.Name
		if fieldName[0] >= 'A' && fieldName[0] <= 'Z' {
			single := getFieldInfo(field)
			result = append(result, single)
		}
	}
	return result
}

func initSqlToArgs(t reflect.Type) sqlToArgsType {
	tKind := getTypeKind(t)
	if tKind == 5 {
		tName := t.String()
		return func(dialect sqlDialect, isInsert bool, v interface{}, in []interface{}, builder *strings.Builder) ([]interface{}, error) {
			return nil, errors.New(fmt.Sprintf("%v dos not support toArgs", tName))
		}
	}

	structToArgs := func(t reflect.Type) func(dialect sqlDialect, isInsert bool, v reflect.Value, in []interface{}, builder *strings.Builder) ([]interface{}, error) {
		fields := getStructPublicField(t)
		return func(dialect sqlDialect, isInsert bool, v reflect.Value, in []interface{}, builder *strings.Builder) ([]interface{}, error) {
			if isInsert == false {
				return nil, errors.New("struct can not to args in none insert ")
			}
			//获取insert的value
			fieldCount := 0
			for _, field := range fields {
				if field.isAutoIncr {
					continue
				}
				if field.isCreated || field.isUpdated {
					in = append(in, time.Now())
				} else {
					if dialect.needToArgsConvert(field.fieldType) {
						//动态数据转换后插入
						argvValue := v.FieldByIndex(field.index)
						in = append(in, dialect.toArgsConvertValue(argvValue))
					} else {
						//直接插入
						in = append(in, v.FieldByIndex(field.index).Interface())
					}

				}
				fieldCount++
			}
			builder.WriteString(getSqlComma(fieldCount))

			return in, nil
		}
	}

	if tKind == 1 {
		structHandler := structToArgs(t)
		return func(dialect sqlDialect, isInsert bool, v interface{}, in []interface{}, builder *strings.Builder) ([]interface{}, error) {
			var err error
			value := reflect.ValueOf(v)
			builder.WriteByte('(')
			in, err = structHandler(dialect, isInsert, value, in, builder)
			if err != nil {
				return nil, err
			}
			builder.WriteByte(')')
			return in, nil
		}
	} else if tKind == 2 {
		structHandler := structToArgs(t.Elem())
		return func(dialect sqlDialect, isInsert bool, v interface{}, in []interface{}, builder *strings.Builder) ([]interface{}, error) {
			var err error
			value := reflect.ValueOf(v).Elem()
			builder.WriteByte('(')
			in, err = structHandler(dialect, isInsert, value, in, builder)
			if err != nil {
				return nil, err
			}
			builder.WriteByte(')')
			return in, nil
		}
	} else {
		structSliceToArgs := func(t reflect.Type) func(dialect sqlDialect, isInsert bool, value reflect.Value, in []interface{}, builder *strings.Builder) ([]interface{}, error) {
			structHandler := structToArgs(t.Elem())

			return func(dialect sqlDialect, isInsert bool, value reflect.Value, in []interface{}, builder *strings.Builder) ([]interface{}, error) {
				var err error
				length := value.Len()
				for i := 0; i != length; i++ {
					if i != 0 {
						builder.WriteString(",(")
					} else {
						builder.WriteByte('(')
					}
					in, err = structHandler(dialect, isInsert, value.Index(i), in, builder)
					if err != nil {
						return nil, err
					}
					builder.WriteByte(')')
				}
				return in, nil
			}
		}

		if tKind == 3 {
			structSliceHandler := structSliceToArgs(t)
			return func(dialect sqlDialect, isInsert bool, v interface{}, in []interface{}, builder *strings.Builder) ([]interface{}, error) {
				value := reflect.ValueOf(v)
				return structSliceHandler(dialect, isInsert, value, in, builder)
			}
		} else {
			structSliceHandler := structSliceToArgs(t.Elem())
			return func(dialect sqlDialect, isInsert bool, v interface{}, in []interface{}, builder *strings.Builder) ([]interface{}, error) {
				value := reflect.ValueOf(v).Elem()
				return structSliceHandler(dialect, isInsert, value, in, builder)
			}
		}
	}
}

func initSqlColumn(t reflect.Type) sqlColumnType {
	tKind := getTypeKind(t)
	if tKind == 5 {
		tName := t.String()
		return func(dialect sqlDialect, isInsert bool, builder *strings.Builder) error {
			return errors.New(fmt.Sprintf("%v dos not support column", tName))
		}
	}

	var fields []sqlStructPublicField
	if tKind == 1 {
		fields = getStructPublicField(t)
	} else if tKind == 2 {
		fields = getStructPublicField(t.Elem())
	} else if tKind == 3 {
		fields = getStructPublicField(t.Elem())
	} else {
		fields = getStructPublicField(t.Elem().Elem())
	}

	return func(dialect sqlDialect, isInsert bool, builder *strings.Builder) error {
		hasData := false
		for _, field := range fields {
			if isInsert && field.isAutoIncr {
				continue
			}
			if hasData {
				builder.WriteByte(',')
			}
			appendQuotedIdentifier(dialect, field.name, builder)
			hasData = true
		}
		return nil
	}
}

func initSqlFromResult(t reflect.Type) sqlFromResultType {
	tKind := getTypeKind(t)
	if tKind != 4 {
		tName := t.String()
		return func(dialect sqlDialect, v interface{}, rows *gosql.Rows) error {
			return errors.New(fmt.Sprintf("%v dos not support fromResult", tName))
		}
	}
	sliceType := t.Elem()
	structType := sliceType.Elem()
	structInfo := getStructPublicField(structType)
	fieldInfoMap := map[string]sqlStructPublicField{}
	for _, single := range structInfo {
		fieldInfoMap[single.name] = single
		fieldInfoMap[strings.ToLower(single.name)] = single
	}
	return func(dialect sqlDialect, v interface{}, rows *gosql.Rows) error {
		//配置列
		columns, err := rows.Columns()
		if err != nil {
			return err
		}
		temp := reflect.New(structType).Elem()
		tempScans := make([]interface{}, len(columns), len(columns))
		tempWriters := []fromResultConvertValueFunc{}
		for i, column := range columns {
			fieldInfo, isExist := fieldInfoMap[column]
			if isExist == false {
				return errors.New(fmt.Sprintf("%v dos not have column %v", structType.String(), column))
			}
			if dialect.needFromResultConvert(fieldInfo.fieldType) {
				//需要动态数据转换
				tempScan, tempWriter := dialect.fromResultConvertValue(temp.FieldByIndex(fieldInfo.index))
				tempScans[i] = tempScan.Addr().Interface()
				if tempWriter != nil {
					tempWriters = append(tempWriters, tempWriter)
				}
			} else {
				//直接转换
				tempScans[i] = temp.FieldByIndex(fieldInfo.index).Addr().Interface()
			}
		}

		//写入数组
		result := reflect.MakeSlice(sliceType, 0, 16)
		for rows.Next() {
			err := rows.Scan(tempScans...)
			if err != nil {
				return err
			}
			for _, tempWriter := range tempWriters {
				tempWriter()
			}
			result = reflect.Append(result, temp)
		}
		reflect.ValueOf(v).Elem().Set(result)
		return nil
	}
}

func initSqlSetValue(t reflect.Type) sqlSetValueType {
	tKind := getTypeKind(t)
	if tKind != 1 && tKind != 2 {
		tName := t.String()
		return func(dialect sqlDialect, v interface{}, in []interface{}, builder *strings.Builder) ([]interface{}, error) {
			return nil, errors.New(fmt.Sprintf("%v dos not support setValue", tName))
		}
	}
	structSetValue := func(t reflect.Type) func(dialect sqlDialect, v reflect.Value, in []interface{}, builder *strings.Builder) ([]interface{}, error) {
		fields := getStructPublicField(t)

		return func(dialect sqlDialect, v reflect.Value, in []interface{}, builder *strings.Builder) ([]interface{}, error) {
			hasData := false
			for _, field := range fields {
				//自增键和created字段不写入
				if field.isAutoIncr == true ||
					field.isCreated == true {
					continue
				}
				if hasData {
					builder.WriteByte(',')
				}

				appendQuotedIdentifier(dialect, field.name, builder)
				builder.WriteString(" = ? ")
				if field.isUpdated == true {
					//updated字段设置为当前时间
					in = append(in, time.Now())
				} else {
					in = append(in, v.FieldByIndex(field.index).Interface())
				}
				hasData = true

			}
			return in, nil
		}
	}

	if tKind == 1 {
		structSetValueHandler := structSetValue(t)
		return func(dialect sqlDialect, v interface{}, in []interface{}, builder *strings.Builder) ([]interface{}, error) {
			value := reflect.ValueOf(v)
			return structSetValueHandler(dialect, value, in, builder)
		}
	} else {
		structSetValueHandler := structSetValue(t.Elem())
		return func(dialect sqlDialect, v interface{}, in []interface{}, builder *strings.Builder) ([]interface{}, error) {
			value := reflect.ValueOf(v).Elem()
			return structSetValueHandler(dialect, value, in, builder)
		}
	}
}

func getSqlComma(num int) string {
	if num < len(commaCache) {
		return commaCache[num]
	} else {
		return strings.Repeat("?,", num-1) + "?"
	}
}

func init() {
	commaCache := make([]string, 128, 128)
	commaCache[1] = "?"
	for i := 2; i != len(commaCache); i++ {
		commaCache[i] = commaCache[i-1] + ",?"
	}
}
