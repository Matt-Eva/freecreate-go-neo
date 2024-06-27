package tests

import (
	"freecreate/internal/utils"
	"strings"
	"testing"
)

func TestBuildFicGenreLabels(t *testing.T){
	type TestCase struct {
		Case []string
		Err bool
		Result string
	}

	cases := []TestCase {
		TestCase {

			[]string{"Horror", "Romance"},
			false,
			":Horror:Romance",
		},
	}

	for _, testCase := range cases {
		labels, err := utils.BuildFicGenreLabel(testCase.Case)
		caseLabels := ":" + strings.Join(testCase.Case, ":")
		if testCase.Err && err == nil {
			t.Errorf("test case labels '%s' are not valid, but no error was thrown", caseLabels)
		}
		if !testCase.Err && err != nil {
			t.Errorf("test case labels '%s' are valid, but no error was thrown", caseLabels)
		}
		if labels != testCase.Result {
			t.Errorf("resulting labels '%s' do not match test case result '%s'", labels, testCase.Result)
		}
	}
}