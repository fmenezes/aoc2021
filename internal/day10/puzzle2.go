package day10

import (
	"sort"
	"strings"
)

func Puzzle2(input string) (int, error) {
	close := map[rune]rune{
		'<': '>',
		'{': '}',
		'[': ']',
		'(': ')',
	}
	points := map[rune]int{
		'>': 4,
		'}': 3,
		']': 2,
		')': 1,
	}

	lines := strings.Split(strings.TrimSpace(input), "\n")
	scores := make([]int, 0, len(lines))
	for _, line := range lines {
		expected := ""
		corrupted := false
		for _, r := range line {
			if isOpenRune(r) {
				expected = string(close[r]) + expected
				continue
			}

			expectedRune := []rune(expected)[0]
			expected = expected[1:]
			if r != expectedRune {
				corrupted = true
			}
		}
		if corrupted {
			continue
		}
		score := 0
		for _, r := range expected {
			score *= 5
			score += points[r]
		}
		scores = append(scores, score)
	}

	sort.Ints(scores)

	middleIndex := len(scores) / 2
	return scores[middleIndex], nil
}
