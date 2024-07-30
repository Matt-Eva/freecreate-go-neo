package utils

import (
	"fmt"
	"testing"
)

func TestMapToStruct(t *testing.T) {
	type TestStruct struct {
		Field  string
		Number int
		MyBool bool
	}
	testStruct := TestStruct{}
	testMap := map[string]any{
		"Field":  "sup",
		"Number": 600,
		"MyBool": true,
	}
	cErr := MapToStruct(testMap, &testStruct)
	if cErr.E != nil {
		cErr.Log()
		t.Fatalf("above error occurred")
	}
	if testStruct.Field != "sup" || testStruct.Number != 600 || testStruct.MyBool != true {
		fmt.Println(testMap, testStruct)
		t.Fatalf("could not convert test map to test struct")
	}
}
