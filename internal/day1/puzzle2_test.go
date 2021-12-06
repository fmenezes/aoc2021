package day1

import "testing"

func Test_Puzzle2(t *testing.T) {
	expected := 5
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
	got, err := Puzzle2(input)
	if err != nil {
		t.Error(err)
	}
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}
