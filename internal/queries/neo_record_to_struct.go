package queries

import (
	"fmt"
	"freecreate/internal/err"
	"reflect"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

func NeoRecordToStruct(record neo4j.Record, structPointer interface{})(err.Error){
	value := reflect.ValueOf(structPointer).Elem()
	t := value.Type()

	for i := 0; i < t.NumField(); i ++{
		field := t.Field(i)
		name := field.Name
		recordValue, ok := record.Get(name)
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
		case reflect.Int:
		case reflect.Int8:
		case reflect.Int16:
		case reflect.Int32:
		case reflect.Int64:
		}
	}

	return err.Error{}
}

func convString(value any, name string)(string, err.Error){
	valid, ok := value.(string)
	if !ok {
		msg := fmt.Sprintf("could not convert value of %s field to string", name)
		return "", err.New(msg)
	}

	return valid, err.Error{}
}