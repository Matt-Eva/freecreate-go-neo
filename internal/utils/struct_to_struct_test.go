package utils

import "testing"

func TestStructToStruct(t *testing.T) {
	type Sender struct {
		StringField string
		NumField    int
		BoolField   bool
	}
	type Receive struct {
		StringField string
		NumField    int
		BoolField   bool
	}

	sender := Sender{
		"string",
		1,
		true,
	}
	var receiver Receive

	cErr := StructToStruct(sender, &receiver)
	if cErr.E != nil {
		cErr.Log()
		t.Fatal("Above error occured")
	}

	if receiver.StringField != sender.StringField {
		t.Errorf("string fields do not match. receiver: %s; sender: %s;", receiver.StringField, sender.StringField)
	}
	if receiver.NumField != sender.NumField {
		t.Errorf("num fields do not match. receiver: %d; sender: %d;", receiver.NumField, sender.NumField)
	}
	if receiver.BoolField != sender.BoolField {
		t.Errorf("bool fields do not match. receiver: %t; sender: %t;", receiver.BoolField, sender.BoolField)

	}

//== testing conversion of struct with fewer fields to struct with more fields ==

	type ShortSender struct {
		StringField string
	}

	type LongReceiver struct {
		StringField string
		NumField int
	}
	shortSender := ShortSender{
		"string",
	}
	var longReceiver LongReceiver
	sErr := StructToStruct(shortSender, &longReceiver)
	if sErr.E != nil {
		sErr.Log()
		t.Fatal("above error occurred")
	}
	
	if longReceiver.StringField != shortSender.StringField {
		t.Error("did not convert matching fields correctly")
	}
	if longReceiver.NumField != 0 {
		t.Error("empty num field not initialized at zero value")
	}
}
