package day2

import (
	"strconv"
	"strings"
)

func Puzzle2(input string) (int, error) {
	horizontal := 0
	depth := 0
	aim := 0
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		columns := strings.SplitN(line, " ", 2)
		direction, sStep := columns[0], columns[1]
		step, err := strconv.Atoi(sStep)
		if err != nil {
			return -1, err
		}
		switch direction {
		case "forward":
			horizontal += step
			depth += (step * aim)
		case "up":
			aim -= step
		case "down":
			aim += step
		}
	}
	return horizontal * depth, nil
}
