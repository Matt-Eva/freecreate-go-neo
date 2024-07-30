package utils

import (
	"reflect"
	"strings"
)

func StructToMap(s interface{}) map[string]any {
	v := reflect.ValueOf(s)
	t := v.Type()
	params := make(map[string]any)

	for i := 0; i < t.NumField(); i++ {
		if v.Field(i).Kind() == reflect.Struct {
			embeddedStruct := v.Field(i)

			for j := 0; j < t.Field(i).Type.NumField(); j++ {
				embeddedFieldName := t.Field(i).Type.Field(j).Name
				embeddedValue := embeddedStruct.Field(j).Interface()
				mapField := strings.ToLower(embeddedFieldName[:1] + embeddedFieldName[1:])
				params[mapField] = embeddedValue
			}
		} else {
			structField := t.Field(i).Name
			mapField := strings.ToLower(structField[:1]) + structField[1:]
			params[mapField] = v.Field(i).Interface()
		}
	}

	return params
}
