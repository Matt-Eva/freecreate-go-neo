package writing

import (
	"testing"
	"time"
)

func TestMakeWriting(t *testing.T) {
	p := PostedWriting{
		Title:       "hello",
		Description: "World",
		WritingType: "test",
		CreatorId:   "1",
	}
	year := time.Now().Year()
	w, wErr := MakeWriting(p, year)
	if wErr.E != nil {
		wErr.Log()
		t.Fatalf("above error occured")
	}
	if w.Title != p.Title {
		t.Errorf("titles do not match")
	}
	if w.Description != p.Description {
		t.Errorf("descriptions do not match")
	}

	if w.CreatorId != p.CreatorId {
		t.Errorf("creator ids do not match")
	}
	if w.OriginalYear != year {
		t.Errorf("years do not match")
	}
	if w.WritingType != p.WritingType {
		t.Errorf("writing types do not match")
	}
}

func TestValidateWriting(t *testing.T) {
	now := time.Now().UnixMilli()
	year := time.Now().Year()
	validWriting := Writing{
		Uid:          "1",
		Title:        "hello",
		Description:  "world",
		Thumbnail:    "",
		WritingType:  "test",
		CreatorId:    "1",
		CreatedAt:    now,
		UpdatedAt:    now,
		OriginalYear: year,
	}
	vErr := validWriting.validateNewWriting(year)
	if vErr.E != nil {
		t.Fatalf("valid writing base case is marked as invalid")
	}

	noUid := validWriting
	noUid.Uid = ""
	uErr := noUid.validateNewWriting(year)
	if uErr.E == nil {
		t.Errorf("validation not catching missing uid")
	}

	noTitle := validWriting
	noTitle.Title = ""
	tErr := noTitle.validateNewWriting(year)
	if tErr.E == nil {
		t.Errorf("validation not catching missing title")
	}

	longTitle := validWriting
	longTitle.Title = "toolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolongtoolong"
	lErr := longTitle.validateNewWriting(year)
	if lErr.E == nil {
		t.Errorf("validation not catching long title")
	}

	longDescription := validWriting
	longDescription.Description = "worldworldworldworldworldworldworldvvworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldvvworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldvvworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldvvworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworldworld"
	dErr := longDescription.validateNewWriting(year)
	if dErr.E == nil {
		t.Errorf("validation not catching long description")
	}

	thumbnail := validWriting
	thumbnail.Thumbnail = "hello"
	thErr := thumbnail.validateNewWriting(year)
	if thErr.E == nil {
		t.Errorf("validation not catching non-empty thumbnail")
	}

	emptyType := validWriting
	emptyType.WritingType = ""
	tyErr := emptyType.validateNewWriting(year)
	if tyErr.E == nil {
		t.Errorf("validation not catching empty type")
	}

	emptyCreator := validWriting
	emptyCreator.CreatorId = ""
	cErr := emptyCreator.validateNewWriting(year)
	if cErr.E == nil {
		t.Errorf("validation not catching empty creator id")
	}

	emptyCreated := validWriting
	emptyCreated.CreatedAt = 0
	caErr := emptyCreated.validateNewWriting(year)
	if caErr.E == nil {
		t.Errorf("validation not catching empty created at")
	}

	emptyUpdated := validWriting
	emptyUpdated.UpdatedAt = 0
	upErr := emptyUpdated.validateNewWriting(year)
	if upErr.E == nil {
		t.Errorf("validation not catching empty updated at")
	}

	nonMatchingCreateAndUpdate := validWriting
	nonMatchingCreateAndUpdate.CreatedAt = 1
	ncuErr := nonMatchingCreateAndUpdate.validateNewWriting(year)
	if ncuErr.E == nil {
		t.Errorf("validation not catching mismatch created and updated at")
	}

	nonZeroLib := validWriting
	nonZeroLib.LibraryCount = 1
	nzlErr := nonZeroLib.validateNewWriting(year)
	if nzlErr.E == nil {
		t.Errorf("validation not catching non zero library count")
	}

	nonZeroLike := validWriting
	nonZeroLike.Likes = 1
	nzliErr := nonZeroLike.validateNewWriting(year)
	if nzliErr.E == nil {
		t.Errorf("validation not catching non zero like count")
	}

	nonZeroDonation := validWriting
	nonZeroDonation.Donations = 1
	nzdErr := nonZeroDonation.validateNewWriting(year)
	if nzdErr.E == nil {
		t.Errorf("validation not catching non zero donation count")
	}

	nonZeroView := validWriting
	nonZeroView.Views = 1
	nzvErr := nonZeroView.validateNewWriting(year)
	if nzvErr.E == nil {
		t.Errorf("validation not catching non zero View count")
	}

	nonZeroRank := validWriting
	nonZeroRank.Rank = 1
	nzrErr := nonZeroRank.validateNewWriting(year)
	if nzrErr.E == nil {
		t.Errorf("validation not catching non zero Rank count")
	}

	nonZeroRelRank := validWriting
	nonZeroRelRank.RelRank = 1
	nzrrErr := nonZeroRelRank.validateNewWriting(year)
	if nzrrErr.E == nil {
		t.Errorf("validation not catching non zero RelRank count")
	}

	nonMatchinYear := validWriting
	nonMatchinYear.OriginalYear = 2000
	yErr := nonMatchinYear.validateNewWriting(year)
	if yErr.E == nil {
		t.Errorf("validation not catching mismatch year")
	}

	published := validWriting
	published.Published = true
	pErr := published.validateNewWriting(year)
	if pErr.E == nil {
		t.Errorf("validation not catching published")
	}

}
