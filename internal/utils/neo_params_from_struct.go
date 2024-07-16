package utils

import (
	"reflect"
	"strings"
)

func NeoParamsFromStruct(s interface{}) map[string]any {
	v := reflect.ValueOf(s)
	t := v.Type()
	params := make(map[string]any)

	for i := 0; i < t.NumField(); i++ {
		structField := t.Field(i).Name
		mapField := strings.ToLower(structField[:1]) + structField[1:]
		params[mapField] = v.Field(i).Interface()
	}

	return params
}
