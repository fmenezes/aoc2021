package day8

import (
	"math"
	"sort"
	"strings"
)

type SortByLen []string

func (s SortByLen) Len() int {
	return len(s)
}

func (s SortByLen) Swap(l, r int) {
	s[l], s[r] = s[r], s[l]
}
func (s SortByLen) Less(l, r int) bool {
	return len(s[l]) < len(s[r])
}

func sortRunes(s string) string {
	l := strings.Split(s, "")
	sort.Strings(l)
	return strings.Join(l, "")
}

func Puzzle2(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	c := 0
	for _, line := range lines {
		columns := strings.Split(line, " | ")
		patterns := strings.Split(strings.TrimSpace(columns[0]), " ")
		sort.Sort(SortByLen(patterns))

		pattern1 := sortRunes(patterns[0]) //len 2
		pattern7 := sortRunes(patterns[1]) //len 3
		pattern4 := sortRunes(patterns[2]) //len 4
		pattern8 := sortRunes(patterns[9]) //len 7

		pattern3 := ""
		pattern5 := ""
		pattern2 := ""
		for _, pattern := range patterns[3:6] { //len 5
			pattern = sortRunes(pattern)
			rc7 := 0
			for _, r := range pattern7 {
				if strings.ContainsRune(pattern, r) {
					rc7 += 1
				}
			}
			if rc7 == 3 {
				pattern3 = pattern
			} else {
				rc4 := 0
				for _, r := range pattern4 {
					if strings.ContainsRune(pattern, r) {
						rc4 += 1
					}
				}
				if rc4 == 3 {
					pattern5 = pattern
				} else if rc4 == 2 {
					pattern2 = pattern
				}
			}
		}

		pattern9 := ""
		pattern0 := ""
		pattern6 := ""
		for _, pattern := range patterns[6:9] { //len 6
			pattern = sortRunes(pattern)
			rc4 := 0
			for _, r := range pattern4 {
				if strings.ContainsRune(pattern, r) {
					rc4 += 1
				}
			}
			if rc4 == 4 {
				pattern9 = pattern
			} else {
				rc1 := 0
				for _, r := range pattern1 {
					if strings.ContainsRune(pattern, r) {
						rc1 += 1
					}
				}
				if rc1 == 2 {
					pattern0 = pattern
				} else {
					pattern6 = pattern
				}
			}
		}

		values := map[string]int{}
		values[pattern1] = 1
		values[pattern2] = 2
		values[pattern3] = 3
		values[pattern4] = 4
		values[pattern5] = 5
		values[pattern6] = 6
		values[pattern7] = 7
		values[pattern8] = 8
		values[pattern9] = 9
		values[pattern0] = 0

		segments := strings.Split(strings.TrimSpace(columns[1]), " ")
		value := 0
		for i, segment := range segments {
			segment = sortRunes(segment)
			factor := len(segments) - i - 1
			factor = int(math.Pow10(factor))
			value += (values[segment] * factor)
		}
		c += value
	}
	return c, nil
}
