package day9

import (
	"strings"
)

func Puzzle1(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	values := make([][]int, len(lines))
	for i, line := range lines {
		values[i] = make([]int, len(line))
		for j, r := range line[:] {
			values[i][j] = int(r - '0')
		}
	}

	risk := 0
	for r, row := range values {
		for c, value := range row {
			lowest := true
			if r > 0 {
				if value >= values[r-1][c] {
					lowest = false
				}
			}
			if r < (len(values) - 1) {
				if value >= values[r+1][c] {
					lowest = false
				}
			}
			if c > 0 {
				if value >= values[r][c-1] {
					lowest = false
				}
			}
			if c < (len(row) - 1) {
				if value >= values[r][c+1] {
					lowest = false
				}
			}
			if lowest {
				risk += value + 1
			}
		}
	}

	return risk, nil
}
