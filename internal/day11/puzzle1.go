package day11

import (
	"fmt"
	"strconv"
	"strings"
)

type myCoordinate struct {
	r, c int
}

func (c myCoordinate) String() string {
	return fmt.Sprintf("%v-%v", c.r, c.c)
}

type myMap struct {
	data [][]int
}

func (m *myMap) increase(c myCoordinate) bool {
	m.data[c.r][c.c] += 1
	if m.data[c.r][c.c] > 9 {
		m.data[c.r][c.c] = 0
		return true
	}
	return false
}

func (m *myMap) step() []myCoordinate {
	flashing := []myCoordinate{}
	for r, row := range m.data {
		for c := range row {
			coord := myCoordinate{r: r, c: c}
			if m.increase(coord) {
				flashing = append(flashing, coord)
			}
		}
	}
	return flashing
}

func (m myMap) adjacents(c myCoordinate) []myCoordinate {
	result := make([]myCoordinate, 0, 8)

	hasRowAbove := c.r > 0
	hasRowBelow := c.r < (len(m.data) - 1)
	hasLeftColumn := c.c > 0
	hasRightColumn := c.c < (len(m.data[c.r]) - 1)

	if hasRowAbove {
		result = append(result, myCoordinate{r: c.r - 1, c: c.c})
	}

	if hasRowBelow {
		result = append(result, myCoordinate{r: c.r + 1, c: c.c})
	}

	if hasLeftColumn {
		result = append(result, myCoordinate{r: c.r, c: c.c - 1})
	}

	if hasRightColumn {
		result = append(result, myCoordinate{r: c.r, c: c.c + 1})
	}

	if hasRowAbove && hasLeftColumn {
		result = append(result, myCoordinate{r: c.r - 1, c: c.c - 1})
	}

	if hasRowAbove && hasRightColumn {
		result = append(result, myCoordinate{r: c.r - 1, c: c.c + 1})
	}

	if hasRowBelow && hasLeftColumn {
		result = append(result, myCoordinate{r: c.r + 1, c: c.c - 1})
	}

	if hasRowBelow && hasRightColumn {
		result = append(result, myCoordinate{r: c.r + 1, c: c.c + 1})
	}

	return result
}

func (m *myMap) increaseRadius(flashingInitially []myCoordinate) int {

	flashing := map[string]int{}
	toVisit := make([]myCoordinate, len(flashingInitially))
	copy(toVisit, flashingInitially)

	for _, e := range flashingInitially {
		flashing[e.String()] = 1
	}

	for len(toVisit) > 0 {
		current := toVisit[0]
		toVisit = toVisit[1:]
		v, ok := flashing[current.String()]
		navigateChildren := false
		if (!ok && m.increase(current)) || v == 1 {
			navigateChildren = true
			flashing[current.String()] = 2
		}
		if navigateChildren {
			toVisit = append(toVisit, m.adjacents(current)...)
		}
	}
	return len(flashing)
}

func (m myMap) String() string {
	r := strings.Builder{}
	for _, row := range m.data {
		for _, column := range row {
			r.WriteString(strconv.Itoa(column))
		}
		r.WriteString("\n")
	}
	return r.String()
}

func Puzzle1(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	values := make([][]int, len(lines))
	for i, line := range lines {
		values[i] = make([]int, len(line))
		for j, r := range line[:] {
			values[i][j] = int(r - '0')
		}
	}

	m := myMap{data: values}
	f := 0
	for d := 0; d < 100; d += 1 {
		flashing := m.step()
		f += m.increaseRadius(flashing)
	}

	return f, nil
}
