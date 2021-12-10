package day5

import (
	"fmt"
	"strconv"
	"strings"
)

func convertCoordinates(s string) (int, int, int, int) {
	sides := strings.SplitN(s, " -> ", 2)
	firstSide := strings.SplitN(sides[0], ",", 2)
	secondSide := strings.SplitN(sides[1], ",", 2)
	x1, _ := strconv.Atoi(firstSide[0])
	y1, _ := strconv.Atoi(firstSide[1])
	x2, _ := strconv.Atoi(secondSide[0])
	y2, _ := strconv.Atoi(secondSide[1])
	return x1, y1, x2, y2
}

func intsSorted(a, b int) (int, int) {
	if a > b {
		return b, a
	}
	return a, b
}

func Puzzle1(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	coordinates := map[int]map[int]int{}
	coordinatesOverlapping := map[string]bool{}

	for _, line := range lines {
		x1, y1, x2, y2 := convertCoordinates(line)
		if x1 == x2 {
			if coordinates[x1] == nil {
				coordinates[x1] = map[int]int{}
			}
			l, r := intsSorted(y1, y2)
			for i := l; i <= r; i += 1 {
				coordinates[x1][i] += 1
				if coordinates[x1][i] > 1 {
					key := fmt.Sprintf("%v-%v", x1, i)
					coordinatesOverlapping[key] = true
				}
			}
		}
		if y1 == y2 {
			l, r := intsSorted(x1, x2)
			for i := l; i <= r; i += 1 {
				if coordinates[i] == nil {
					coordinates[i] = map[int]int{}
				}
				coordinates[i][y1] += 1
				if coordinates[i][y1] > 1 {
					key := fmt.Sprintf("%v-%v", i, y1)
					coordinatesOverlapping[key] = true
				}
			}
		}
	}

	return len(coordinatesOverlapping), nil
}
