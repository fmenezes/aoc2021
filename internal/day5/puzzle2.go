package day5

import (
	"strings"
)

type coordinate struct {
	x, y int
}

type counter struct {
	coordinates            map[int]map[int]int
	coordinatesOverlapping map[coordinate]bool
}

func newCounter() *counter {
	return &counter{coordinates: map[int]map[int]int{}, coordinatesOverlapping: map[coordinate]bool{}}
}

func (c *counter) countCoordinates(co coordinate) {
	if c.coordinates == nil {
		c.coordinates = map[int]map[int]int{}
	}
	if c.coordinates[co.x] == nil {
		c.coordinates[co.x] = map[int]int{}
	}
	c.coordinates[co.x][co.y] += 1
	if c.coordinates[co.x][co.y] > 1 {
		c.coordinatesOverlapping[co] = true
	}
}

func (c *counter) generateCoordinates(input string) chan coordinate {
	ch := make(chan coordinate)
	go func() {
		lines := strings.Split(strings.TrimSpace(input), "\n")
		for _, line := range lines {
			x1, y1, x2, y2 := convertCoordinates(line)
			if x1 == x2 {
				l, r := intsSorted(y1, y2)
				for i := l; i <= r; i += 1 {
					ch <- coordinate{x: x1, y: i}
				}
			} else if y1 == y2 {
				l, r := intsSorted(x1, x2)
				for i := l; i <= r; i += 1 {
					ch <- coordinate{x: i, y: y1}
				}
			} else {
				lx, rx := intsSorted(x1, x2)
				ly, ry := 0, 0

				if lx != x1 {
					ly, ry = y2, y1
				} else {
					ly, ry = y1, y2
				}
				y := ly
				for x := lx; x <= rx; x += 1 {
					ch <- coordinate{x: x, y: y}
					if ly > ry {
						y -= 1
					} else {
						y += 1
					}
				}

			}
		}
		close(ch)
	}()

	return ch
}

func Puzzle2(input string) (int, error) {

	c := newCounter()

	for co := range c.generateCoordinates(input) {
		c.countCoordinates(co)
	}

	return len(c.coordinatesOverlapping), nil
}
