package day7

import "testing"

func Test_Puzzle2(t *testing.T) {
	expected := 168
	input := `16,1,2,0,4,2,7,1,2,14`
	got, err := Puzzle2(input)
	if err != nil {
		t.Error(err)
	}
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}
