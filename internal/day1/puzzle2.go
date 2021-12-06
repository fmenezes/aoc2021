package day1

import (
	"strconv"
	"strings"
)

func Puzzle2(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	c := 0
	lastSum := -1
	count := []int{0, -1, -2}
	sum := []int{0, 0, 0}

	for i := 0; i < len(lines); i += 1 {
		line := lines[i]
		intLine, err := strconv.Atoi(line)
		if err != nil {
			return -1, err
		}

		for j := 0; j < len(count); j += 1 {
			count[j] += 1
			if count[j] > 0 {
				sum[j] += intLine
			}

			if count[j] == len(count) {
				if lastSum > 0 && sum[j] > lastSum {
					c += 1
				}
				lastSum = sum[j]
				count[j] = 0
				sum[j] = 0
			}

		}
	}
	return c, nil
}
