package queries

import (
	"fmt"
	"testing"
)

func TestNeoParamsFromStruct(t *testing.T) {
	type Embedded struct {
		Inner string
	}
	type Outer struct {
		Embedded Embedded
		Outer    string
	}
	testStruct := Outer{
		Embedded{
			"hello",
		},
		"world",
	}
	fmt.Println(testStruct.Embedded.Inner)
	testMap := map[string]any{
		"inner": "hello",
		"outer": "world",
	}
	generatedMap := NeoParamsFromStruct(testStruct)
	for key, val := range generatedMap {
		testVal, ok := testMap[key]
		if !ok {
			t.Errorf("key '%s' from generated map does not exist it test map", key)
		}

		if val != testVal {
			t.Errorf("value from generated map '%s' does not match value from test map '%s'", val, testVal)
		}
	}
}
