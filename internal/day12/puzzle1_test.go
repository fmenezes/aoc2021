package day12

import "testing"

func Test_Puzzle1(t *testing.T) {
	expected := 10
	input := `start-A
start-b
A-c
A-b
b-d
A-end
b-end`
	got, err := Puzzle1(input)
	if err != nil {
		t.Error(err)
	}
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}
