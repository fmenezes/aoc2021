package day7

import (
	"strconv"
	"strings"
)

func factorial(n int) int {
	if n <= 0 {
		return 0
	} else {
		return n + factorial(n-1)
	}
}

func Puzzle2(input string) (int, error) {
	values := strings.Split(strings.TrimSpace(input), ",")
	nums := make([]int, len(values))
	distances := map[int]int{}
	for i, v := range values {
		n, err := strconv.Atoi(v)
		if err != nil {
			return 0, err
		}
		nums[i] = n
	}
	win := -1
	for _, n := range nums {
		if distances[n] <= 0 {
			for _, i := range nums {
				d := i - n
				if d < 0 {
					d = d * -1
				}
				distances[n] += factorial(d)
			}
		}
		if win == -1 || distances[n] < win {
			win = distances[n]
		}
	}

	return win, nil
}
