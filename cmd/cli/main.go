package main

import (
	"errors"
	"flag"
	"fmt"
	"log"

	"github.com/fmenezes/aoc2021/internal"
	"github.com/fmenezes/aoc2021/internal/day1"
	"github.com/fmenezes/aoc2021/internal/day10"
	"github.com/fmenezes/aoc2021/internal/day11"
	"github.com/fmenezes/aoc2021/internal/day12"
	"github.com/fmenezes/aoc2021/internal/day2"
	"github.com/fmenezes/aoc2021/internal/day3"
	"github.com/fmenezes/aoc2021/internal/day4"
	"github.com/fmenezes/aoc2021/internal/day5"
	"github.com/fmenezes/aoc2021/internal/day6"
	"github.com/fmenezes/aoc2021/internal/day7"
	"github.com/fmenezes/aoc2021/internal/day8"
	"github.com/fmenezes/aoc2021/internal/day9"
)

func main() {
	var day, puzzle int
	flag.IntVar(&day, "d", 0, "day")
	flag.IntVar(&puzzle, "p", 0, "puzzle")

	flag.Parse()

	listOfPuzzles := map[string]func(string) (int, error){
		"d1p1":  day1.Puzzle1,
		"d1p2":  day1.Puzzle2,
		"d2p1":  day2.Puzzle1,
		"d2p2":  day2.Puzzle2,
		"d3p1":  day3.Puzzle1,
		"d3p2":  day3.Puzzle2,
		"d4p1":  day4.Puzzle1,
		"d4p2":  day4.Puzzle2,
		"d5p1":  day5.Puzzle1,
		"d5p2":  day5.Puzzle2,
		"d6p1":  day6.Puzzle1,
		"d6p2":  day6.Puzzle2,
		"d7p1":  day7.Puzzle1,
		"d7p2":  day7.Puzzle2,
		"d8p1":  day8.Puzzle1,
		"d8p2":  day8.Puzzle2,
		"d9p1":  day9.Puzzle1,
		"d9p2":  day9.Puzzle2,
		"d10p1": day10.Puzzle1,
		"d10p2": day10.Puzzle2,
		"d11p1": day11.Puzzle1,
		"d11p2": day11.Puzzle2,
		"d12p1": day12.Puzzle1,
	}

	if day == 0 || puzzle == 0 {
		day = 12
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
