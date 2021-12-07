package day4

import (
	"strconv"
	"strings"
)

func Puzzle2(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	firstLine := strings.Split(lines[0], ",")
	turns := map[string]int{}
	for i, entry := range firstLine {
		turns[entry] = i
	}
	board := 0
	lastToWinBoard := 0
	lastToWinBoardTurn := -1
	turnPerBoard := -1
	turnPerColumn := []int{-1, -1, -1, -1, -1}
	for _, line := range lines[2:] {
		if len(line) == 0 {
			for _, turn := range turnPerColumn {
				if turn < turnPerBoard {
					turnPerBoard = turn
				}
			}
			if lastToWinBoardTurn == -1 || turnPerBoard > lastToWinBoardTurn {
				lastToWinBoard = board
				lastToWinBoardTurn = turnPerBoard
			}
			board += 1
			turnPerBoard = -1
			turnPerColumn = []int{-1, -1, -1, -1, -1}
			continue
		}
		turnPerLine := -1
		for c, i := 0, 0; c < len(line); c, i = c+3, i+1 {
			item := strings.TrimSpace(line[c : c+2])
			if turns[item] > turnPerColumn[i] {
				turnPerColumn[i] = turns[item]
			}
			if turns[item] > turnPerLine {
				turnPerLine = turns[item]
			}
		}
		if turnPerBoard == -1 || turnPerLine < turnPerBoard {
			turnPerBoard = turnPerLine
		}
	}
	for _, turn := range turnPerColumn {
		if turn < turnPerBoard {
			turnPerBoard = turn
		}
	}
	if turnPerBoard > lastToWinBoardTurn {
		lastToWinBoard = board
		lastToWinBoardTurn = turnPerBoard
	}

	l := 2 + (lastToWinBoard * 6)
	sumOfUnmarked := 0
	latestTurn := -1
	for _, line := range lines[l : l+5] {
		for c, i := 0, 0; c < len(line); c, i = c+3, i+1 {
			item := strings.TrimSpace(line[c : c+2])
			v, err := strconv.Atoi(item)
			if err != nil {
				return 0, err
			}
			if turns[item] == lastToWinBoardTurn {
				latestTurn = v
			}
			if turns[item] > lastToWinBoardTurn {
				sumOfUnmarked += v
			}
		}
	}
	score := latestTurn * sumOfUnmarked
	return score, nil
}
