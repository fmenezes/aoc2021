package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	d1p1 "github.com/fmenezes/aoc2021/internal/day1/puzzle1"
	d1p2 "github.com/fmenezes/aoc2021/internal/day1/puzzle2"
	d2p1 "github.com/fmenezes/aoc2021/internal/day2/puzzle1"
	d2p2 "github.com/fmenezes/aoc2021/internal/day2/puzzle2"
)

func main() {
	var day, puzzle int
	flag.IntVar(&day, "d", 0, "day")
	flag.IntVar(&puzzle, "p", 0, "puzzle")

	flag.Parse()

	listOfPuzzles := map[string]func() error{
		"d1p1": d1p1.Run,
		"d1p2": d1p2.Run,
		"d2p1": d2p1.Run,
		"d2p2": d2p2.Run,
	}

	if day == 0 || puzzle == 0 {
		day = 2
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
