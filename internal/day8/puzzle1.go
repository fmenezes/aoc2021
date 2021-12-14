package day8

import (
	"strings"
)

func Puzzle1(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	c := 0
	for _, line := range lines {
		columns := strings.Split(line, " | ")
		segments := strings.Split(strings.TrimSpace(columns[1]), " ")
		for _, segment := range segments {
			if len(segment) == 2 || len(segment) == 3 || len(segment) == 4 || len(segment) == 7 {
				c += 1
			}
		}
	}
	return c, nil
}
