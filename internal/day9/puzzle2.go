package day9

import (
	"fmt"
	"sort"
	"strings"
)

type coordinates struct {
	r, c int
}

type myMap struct {
	values [][]int
}

func (c coordinates) String() string {
	return fmt.Sprintf("%v-%v", c.r, c.c)
}
func (m myMap) Value(c coordinates) int {
	return m.values[c.r][c.c]
}

func (m myMap) basin(c coordinates) int {
	toVisit := []coordinates{c}
	compare := m.Value(c)
	visited := map[string]bool{}
	count := 0
	for len(toVisit) > 0 {
		current := toVisit[0]
		toVisit = toVisit[1:]
		if visited[current.String()] {
			continue
		}
		visited[current.String()] = true
		if m.Value(current) < 9 && m.Value(current) >= compare {
			count += 1
			toVisit = append(toVisit, m.adjacents(current)...)
		}
	}
	return count
}

func (m myMap) adjacents(c coordinates) []coordinates {
	result := make([]coordinates, 0, 4)

	if c.r > 0 {
		result = append(result, coordinates{r: c.r - 1, c: c.c})
	}

	if c.r < (len(m.values) - 1) {
		result = append(result, coordinates{r: c.r + 1, c: c.c})
	}

	if c.c > 0 {
		result = append(result, coordinates{r: c.r, c: c.c - 1})
	}

	if c.c < (len(m.values[c.r]) - 1) {
		result = append(result, coordinates{r: c.r, c: c.c + 1})
	}

	return result
}

func (m myMap) lowest(c coordinates) bool {
	lowest := true
	value := m.Value(c)
	for _, c := range m.adjacents(c) {
		if value >= m.Value(c) {
			lowest = false
		}
	}
	return lowest
}

func Puzzle2(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")
	values := make([][]int, len(lines))
	for i, line := range lines {
		values[i] = make([]int, len(line))
		for j, r := range line[:] {
			values[i][j] = int(r - '0')
		}
	}

	m := myMap{values: values}

	basins := []int{}
	for r, row := range m.values {
		for c := range row {
			coordinate := coordinates{r: r, c: c}
			if m.lowest(coordinate) {
				basins = append(basins, m.basin(coordinate))
			}
		}
	}

	if len(basins) > 3 {
		sort.Ints(basins)
		basins = basins[len(basins)-3:]
	}

	res := 1
	for _, basin := range basins {
		res *= basin
	}

	return res, nil
}
