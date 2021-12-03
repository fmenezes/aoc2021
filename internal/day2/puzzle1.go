package day2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fmenezes/aoc2021/internal"
)

func Puzzle1() error {
	input, err := internal.Input(2)
	if err != nil {
		return err
	}
	horizontal := 0
	depth := 0
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for _, line := range lines {
		columns := strings.SplitN(line, " ", 2)
		direction, sStep := columns[0], columns[1]
		step, err := strconv.Atoi(sStep)
		if err != nil {
			return err
		}
		switch direction {
		case "forward":
			horizontal += step
		case "up":
			depth -= step
		case "down":
			depth += step
		}
	}
	fmt.Printf("%v\n", horizontal*depth)
	return nil
}
