package day6

import (
	"strconv"
	"strings"
)

func Puzzle1(input string) (int, error) {
	data := []int{}

	values := strings.Split(strings.TrimSpace(input), ",")
	for _, v := range values {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
		data = append(data, i)
	}

	for d := 0; d < 80; d += 1 {
		for i := 0; i < len(data); i += 1 {
			data[i] -= 1
			if data[i] == -1 {
				data[i] = 6
				data = append(data, 9)
			}
		}

	}

	return len(data), nil
}
