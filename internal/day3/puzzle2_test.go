package day3

import "testing"

func Test_Puzzle2(t *testing.T) {
	expected := 230
	input := `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`
	got, err := Puzzle2(input)
	if err != nil {
		t.Error(err)
	}
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}
