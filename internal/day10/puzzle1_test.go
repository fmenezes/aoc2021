package day10

import "testing"

func Test_Puzzle1(t *testing.T) {
	expected := 26397
	input := `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`
	got, err := Puzzle1(input)
	if err != nil {
		t.Error(err)
	}
	if got != expected {
		t.Errorf("expected: %v, got: %v", expected, got)
	}
}
