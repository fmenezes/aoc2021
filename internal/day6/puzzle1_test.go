package day6

import "testing"

func Test_Puzzle1(t *testing.T) {
	expected := 5934
	input := `3,4,3,1,2`
	got, err := Puzzle1(input)
	if err != nil {
		t.Error(err)
	}
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}
