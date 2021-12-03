package puzzle2

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fmenezes/aoc2021/internal"
)

func Run() error {
	input, err := internal.Input(1)
	if err != nil {
		return err
	}

	lines := strings.Split(strings.TrimSpace(input), "\n")

	c := 0
	lastSum := -1
	count := []int{0, -1, -2}
	sum := []int{0, 0, 0}

	for i := 0; i < len(lines); i += 1 {
		line := lines[i]
		intLine, err := strconv.Atoi(line)
		if err != nil {
			return err
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
	fmt.Println(c)
	return nil
}
