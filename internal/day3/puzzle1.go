package day3

import (
	"strings"
)

func convertLine(line string) []int {
	runes := []rune(line)
	result := make([]int, len(runes))
	for i, r := range runes {
		result[i] = int(r - '0')
	}
	return result
}

func Puzzle1(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	sums := convertLine(lines[0])
	for _, line := range lines[1:] {
		bits := convertLine(line)
		for i, bit := range bits {
			sums[i] += bit
		}
	}
	gammaRate := 0
	epsilonRate := 0
	for _, sum := range sums {
		currentGammaRateBit := 0
		currentEpsilonRateBit := 1
		if sum >= (len(lines) / 2) {
			currentGammaRateBit = 1
			currentEpsilonRateBit = 0
		}
		gammaRate = gammaRate << 1
		gammaRate = gammaRate | currentGammaRateBit
		epsilonRate = epsilonRate << 1
		epsilonRate = epsilonRate | currentEpsilonRateBit
	}

	return gammaRate * epsilonRate, nil
}
