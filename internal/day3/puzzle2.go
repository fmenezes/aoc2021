package day3

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

type byBits struct {
	data     []string
	bitLen   int
	position int
}

func (v byBits) Len() int {
	return len(v.data)
}

func (v byBits) Swap(i, j int) {
	v.data[i], v.data[j] = v.data[j], v.data[i]
}

func (v byBits) Less(i, j int) bool {
	return v.getBit(i) < v.getBit(j)
}

func (v byBits) getBit(row int) int {
	b, _ := strconv.Atoi(v.data[row][v.position : v.position+1])
	return b
}

func (v byBits) search(most bool) string {
	c := byBits{
		bitLen:   v.bitLen,
		data:     v.data[:],
		position: 0,
	}
	for c.Len() > 1 {
		p := sort.Search(c.Len(), func(i int) bool {
			return c.getBit(i) == 1
		})
		var part []string
		half := int(math.Floor(float64(c.Len()) / 2))

		keepZeros := (most && p > half) || (!most && p <= half)
		if keepZeros {
			part = c.data[:p]
		} else {
			part = c.data[p:]
		}
		c = byBits{
			bitLen:   c.bitLen,
			data:     part,
			position: c.position + 1,
		}
	}
	return c.data[0]
}

func Puzzle2(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	sort.Strings(lines)
	bits := byBits{
		bitLen: len(lines[0]),
		data:   lines,
	}
	most := bits.search(true)
	less := bits.search(false)

	m, err := strconv.ParseInt(most, 2, 0)
	if err != nil {
		return -1, err
	}

	l, err := strconv.ParseInt(less, 2, 0)
	if err != nil {
		return -1, err
	}

	return int(m * l), nil
}
