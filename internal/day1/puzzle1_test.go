package day1

import "testing"

func Test_Puzzle1(t *testing.T) {
	expected := 7
	input := `199
200
208
210
200
207
240
269
260
263`
	got, err := Puzzle1(input)
	if err != nil {
		t.Error(err)
	}
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}
