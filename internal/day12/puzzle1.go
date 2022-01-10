package day12

import (
	"strings"
	"unicode"
)

type graph struct {
	edges map[string][]string
}

func (g *graph) addNode(dst, src string) {
	if g.edges[src] == nil {
		g.edges[src] = []string{dst}
	} else {
		g.edges[src] = append(g.edges[src], dst)
	}

	if g.edges[dst] == nil {
		g.edges[dst] = []string{src}
	} else {
		g.edges[dst] = append(g.edges[dst], src)
	}
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func (g *graph) traversePaths(startNode string) chan []string {
	c := make(chan []string)

	go func() {
		toVisit := []string{startNode}
		visited := map[string]bool{}
		for len(toVisit) > 0 {
			current := toVisit[0]
			toVisit = toVisit[1:]
			if visited[current] {
				continue
			}
			visited[current] = true
			fullPath := strings.Split(current, ",")
			c <- fullPath
			alreadyInPath := map[string]bool{}
			for _, entry := range fullPath {
				alreadyInPath[entry] = true
			}
			lastPath := fullPath[len(fullPath)-1]
			children := g.edges[lastPath]
			for _, child := range children {
				if IsUpper(child) || !alreadyInPath[child] {
					toVisit = append(toVisit, strings.Join([]string{current, child}, ","))
				}
			}
		}
		close(c)
	}()

	return c
}

func Puzzle1(input string) (int, error) {
	lines := strings.Split(strings.TrimSpace(input), "\n")

	g := graph{edges: make(map[string][]string, len(lines))}

	for _, line := range lines {
		edges := strings.SplitN(line, "-", 2)
		g.addNode(edges[0], edges[1])
	}

	c := 0
	for fullPath := range g.traversePaths("start") {
		lastPath := fullPath[len(fullPath)-1]
		if lastPath == "end" {
			c += 1
		}
	}

	return c, nil
}
