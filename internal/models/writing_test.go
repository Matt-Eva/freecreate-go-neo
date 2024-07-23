package models

import (
	"testing"
	"time"
)

func TestMakeWriting(t *testing.T){
	p := PostedWriting{
		Title: "hello",
		Description: "World",
		WritingType: "test",
		Thumbnail: "",
		CreatorId: "1",
	}
	year := time.Now().Year()
	w, wErr := MakeWriting(p, year)
	if wErr.E != nil {
		wErr.Log()
		t.Fatalf("above error occured")
	}
	if w.Title != p.Title{
		t.Fatalf("titles do not match")
	}
	if w.Description != p.Description {
		t.Fatalf("descriptions do not match")
	}
	if w.Thumbnail != p.Thumbnail{
		t.Fatalf("thumbnails do not match")
	}
	if w.CreatorId != p.CreatorId{
		t.Fatalf("creator ids do not match")
	}
	if w.OriginalYear != year{
		t.Fatalf("years do not match")
	}
	if w.WritingType != p.WritingType{
		t.Fatalf("writing types do not match")
	}
}