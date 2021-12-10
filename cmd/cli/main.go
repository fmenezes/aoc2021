package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/fmenezes/aoc2021/internal"
	"github.com/fmenezes/aoc2021/internal/day1"
	"github.com/fmenezes/aoc2021/internal/day2"
	"github.com/fmenezes/aoc2021/internal/day3"
	"github.com/fmenezes/aoc2021/internal/day4"
	"github.com/fmenezes/aoc2021/internal/day5"
	"github.com/fmenezes/aoc2021/internal/day6"
)

func main() {
	var day, puzzle int
	flag.IntVar(&day, "d", 0, "day")
	flag.IntVar(&puzzle, "p", 0, "puzzle")

	flag.Parse()

	listOfPuzzles := map[string]func(string) (int, error){
		"d1p1": day1.Puzzle1,
		"d1p2": day1.Puzzle2,
		"d2p1": day2.Puzzle1,
		"d2p2": day2.Puzzle2,
		"d3p1": day3.Puzzle1,
		"d3p2": day3.Puzzle2,
		"d4p1": day4.Puzzle1,
		"d4p2": day4.Puzzle2,
		"d5p1": day5.Puzzle1,
		"d5p2": day5.Puzzle2,
		"d6p1": day6.Puzzle1,
	}

	if day == 0 || puzzle == 0 {
		day = 6
		puzzle = 1
	}

	sel := fmt.Sprintf("d%vp%v", day, puzzle)

	puzzleFunc := listOfPuzzles[sel]

	if puzzleFunc == nil {
		log.Panicln(errors.New("puzzle not found"))
	}

	input, err := internal.Input(day)
	if err != nil {
		log.Panicln(err)
	}

	r, err := puzzleFunc(input)
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(r)
}
