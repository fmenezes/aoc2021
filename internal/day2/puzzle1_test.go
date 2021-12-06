package day2

import "testing"

func Test_Puzzle1(t *testing.T) {
	expected := 150
	input := `forward 5
down 5
forward 8
up 3
down 8
forward 2`
	got, err := Puzzle1(input)
	if err != nil {
		t.Error(err)
	}
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}
