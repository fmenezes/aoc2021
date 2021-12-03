package puzzle2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fmenezes/aoc2021/inputs"
)

func Run() error {
	input, err := inputs.Input(2)
	if err != nil {
		return err
	}
	horizontal := 0
	depth := 0
	aim := 0
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
			depth += (step * aim)
		case "up":
			aim -= step
		case "down":
			aim += step
		}
	}
	fmt.Printf("%v\n", horizontal*depth)
	return nil
}
