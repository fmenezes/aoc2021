package day6

import "testing"

func Test_Puzzle2(t *testing.T) {
	expected := 26984457539
	input := `3,4,3,1,2`
	got, err := Puzzle2(input)
	if err != nil {
		t.Error(err)
	}
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}
