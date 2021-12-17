package day10

import "strings"

func isOpenRune(r rune) bool {
	return r == '<' || r == '{' || r == '[' || r == '('
}

func Puzzle1(input string) (int, error) {
	close := map[rune]rune{
		'<': '>',
		'{': '}',
		'[': ']',
		'(': ')',
	}
	points := map[rune]int{
		'>': 25137,
		'}': 1197,
		']': 57,
		')': 3,
	}

	lines := strings.Split(strings.TrimSpace(input), "\n")
	counts := map[rune]int{}
	for _, line := range lines {
		expected := ""
		for _, r := range line {
			if isOpenRune(r) {
				expected = string(close[r]) + expected
				continue

			}
			expectedRune := []rune(expected)[0]
			expected = expected[1:]
			if r != expectedRune {
				counts[r] += 1
			}
		}
	}

	sum := 0
	for r, c := range counts {
		sum += (points[r] * c)
	}

	return sum, nil
}
