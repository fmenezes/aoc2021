package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	d1p1 "github.com/fmenezes/adventofcode2021/day1/puzzle1"
	d1p2 "github.com/fmenezes/adventofcode2021/day1/puzzle2"
)

func main() {
	var day, puzzle int
	flag.IntVar(&day, "d", 0, "day")
	flag.IntVar(&puzzle, "p", 0, "puzzle")

	flag.Parse()

	listOfPuzzles := map[string]func() error{"d1p1": d1p1.Run, "d1p2": d1p2.Run}

	if day == 0 || puzzle == 0 {
		day = 1
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
