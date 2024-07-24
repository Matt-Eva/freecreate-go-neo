package queries

import "testing"

func TestBuildGenreLabels(t *testing.T){
	genres := []string{"Fantasy", "ScienceFiction","Action"}
	validLabels := ":Action:Fantasy:ScienceFiction"
	generatedLabels, gErr := BuildGenreLabels(genres)
	if gErr.E !=nil {
		gErr.Log()
		t.Fatal("above error occurred")
	}
	if validLabels != generatedLabels{
		t.Fatalf("valid labels %s do not match generated labels %s", validLabels, generatedLabels)
	}
}