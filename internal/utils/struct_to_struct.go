package utils

import (
	"fmt"
	"freecreate/internal/err"
	"reflect"
)

func StructToStruct(sender interface{}, receiverPntr interface{}) err.Error {
	receiverValue := reflect.ValueOf(receiverPntr).Elem()
	receiverType := receiverValue.Type()

	senderValue := reflect.ValueOf(sender)

	for i := 0; i < receiverType.NumField(); i++ {
		field := receiverType.Field(i)
		name := field.Name
		receiverFieldType := field.Type.Kind()
		receiverFieldValue := receiverValue.Field(i)

		senderFieldValue := senderValue.FieldByName(name)
		if !senderFieldValue.IsValid() {
			continue
		}
		senderFieldType := senderFieldValue.Type().Kind()

		if receiverFieldType != senderFieldType {
			msg := fmt.Sprintf("struct field %s field type does not match", name)
			return err.New(msg)
		}
		receiverFieldValue.Set(senderFieldValue)
	}

	return err.Error{}
}
