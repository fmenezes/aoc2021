package day1

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/fmenezes/aoc2021/internal"
)

func Puzzle1() error {
	input, err := internal.Input(1)
	if err != nil {
		return err
	}
	c := 0
	lines := strings.Split(strings.TrimSpace(input), "\n")
	for i := 1; i < len(lines); i += 1 {
		prevLine := lines[i-1]
		line := lines[i]
		intLine, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		intPrevLine, err := strconv.Atoi(prevLine)
		if err != nil {
			return err
		}
		if intLine > intPrevLine {
			c += 1
		}
	}
	fmt.Println(c)
	return nil
}
