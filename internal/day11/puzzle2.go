package day11

import (
	"strings"
)

func Puzzle2(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	values := make([][]int, len(lines))
	for i, line := range lines {
		values[i] = make([]int, len(line))
		for j, r := range line[:] {
			values[i][j] = int(r - '0')
		}
	}

	m := myMap{data: values}
	for d := 0; ; d += 1 {
		flashing := m.step()
		f := m.increaseRadius(flashing)
		if f == 100 {
			return d + 1, nil
		}
	}

	return -1, nil
}
