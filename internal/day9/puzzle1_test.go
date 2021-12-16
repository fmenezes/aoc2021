package day9

import "testing"

func Test_Puzzle1(t *testing.T) {
	expected := 15
	input := `2199943210
3987894921
9856789892
8767896789
9899965678`
	got, err := Puzzle1(input)
	if err != nil {
		t.Error(err)
	}
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}
