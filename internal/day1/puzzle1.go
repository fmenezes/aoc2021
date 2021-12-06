package day1

import (
	"strconv"
	"strings"
)

func Puzzle1(input string) (int, error) {
	c := 0
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for i := 1; i < len(lines); i += 1 {
		prevLine := lines[i-1]
		line := lines[i]
		intLine, err := strconv.Atoi(line)
		if err != nil {
			return -1, err
		}
		intPrevLine, err := strconv.Atoi(prevLine)
		if err != nil {
			return -1, err
		}
		if intLine > intPrevLine {
			c += 1
		}
	}
	return c, nil
}
