package day6

import (
	"strconv"
	"strings"
)

func Puzzle2(input string) (int, error) {
	data := []int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	values := strings.Split(strings.TrimSpace(input), ",")
	for _, v := range values {
		i, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
		data[i] += 1
	}

	sum := 0
	for d := 0; d < 256; d += 1 {
		temp := data[0]
		sum = 0
		for i := 1; i < len(data); i += 1 {
			data[i-1] = data[i]
			sum += data[i]
		}
		data[8] = temp
		data[6] += temp
		sum += (temp * 2)
	}

	return sum, nil
}
