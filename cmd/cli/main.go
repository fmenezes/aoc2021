package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/fmenezes/aoc2021/internal/day1"
	"github.com/fmenezes/aoc2021/internal/day2"
	"github.com/fmenezes/aoc2021/internal/day3"
)

func main() {
	var day, puzzle int
	flag.IntVar(&day, "d", 0, "day")
	flag.IntVar(&puzzle, "p", 0, "puzzle")

	flag.Parse()

	listOfPuzzles := map[string]func() error{
		"d1p1": day1.Puzzle1,
		"d1p2": day1.Puzzle2,
		"d2p1": day2.Puzzle1,
		"d2p2": day2.Puzzle2,
		"d3p1": day3.Puzzle1,
		"d3p2": day3.Puzzle2,
	}

	if day == 0 || puzzle == 0 {
		day = 3
		puzzle = 2
	}

	sel := fmt.Sprintf("d%vp%v", day, puzzle)

	puzzleFunc := listOfPuzzles[sel]

	if puzzleFunc == nil {
		log.Panicln(errors.New("puzzle not found"))
	}

	if err := puzzleFunc(); err != nil {
		log.Panicln(err)
	}
}
