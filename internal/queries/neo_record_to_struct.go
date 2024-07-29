package queries

import (
	"fmt"
	"freecreate/internal/err"
	"reflect"
)

func NeoRecordToStruct(record map[string]any, structPointer interface{}) err.Error {
	value := reflect.ValueOf(structPointer).Elem()
	t := value.Type()

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		name := field.Name
		recordValue, ok := record[name]
		if !ok {
			msg := fmt.Sprintf("neo4j record does not have attribute %s", name)
			return err.New(msg)
		}

		fieldType := field.Type.Kind()
		fieldValue := value.Field(i)
		switch fieldType {
		case reflect.String:
			valid, sErr := convString(recordValue, name)
			if sErr.E != nil {
				return sErr
			}
			fieldValue.SetString(valid)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			valid, iErr := convInt(recordValue, name)
			if iErr.E != nil {
				return iErr
			}
			fieldValue.SetInt(valid)
		case reflect.Bool:
			valid, iErr := convBool(recordValue, name)
			if iErr.E != nil {
				return iErr
			}
			fieldValue.SetBool(valid)
		default:
			return err.New("field value not of type int, bool, or string")
		}
	}

	return err.Error{}
}

func convString(value interface{}, name string) (string, err.Error) {
	switch v := value.(type) {
	case string:
		return v, err.Error{}
	default:
		msg := fmt.Sprintf("record field for '%s' is not a string", name)
		return "", err.New(msg)
	}
}

func convInt(value interface{}, name string) (int64, err.Error) {
	switch v := value.(type) {
	case int64:
		return v, err.Error{}
	case int:
		return int64(v), err.Error{}
	case int8:
		return int64(v), err.Error{}
	case int16:
		return int64(v), err.Error{}
	case int32:
		return int64(v), err.Error{}
	default:
		msg := fmt.Sprintf("map value for field %s is not an integer", name)
		return 0, err.New(msg)
	}
}

func convBool(value interface{}, name string) (bool, err.Error) {
	switch v := value.(type) {
	case bool:
		return v, err.Error{}
	default:
		msg := fmt.Sprintf("record field for '%s' is not a boolean", name)
		return false, err.New(msg)
	}
}
