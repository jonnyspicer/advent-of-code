package day

import (
	"strings"
	"unicode"
)

func Twelve(input []string) (int, int) {
	caves := map[string]*node{}

	for _, row := range input {
		lr := strings.Split(row, "-")
		l, r := lr[0], strings.TrimSuffix(lr[1], "\r")

		if _, ok := caves[l]; ok {
			// if entry for cave exists in map, update it
			caves[l].vertices = append(caves[l].vertices, r)
		} else {
			caves[l] = &node{
				name:     l,
				vertices: []string{r},
				big:      unicode.IsUpper(rune(l[0])),
			}
		}

		if _, ok := caves[r]; ok {
			// if entry for cave exists in map, update it
			caves[r].vertices = append(caves[r].vertices, l)
		} else {
			caves[r] = &node{
				name:     r,
				vertices: []string{l},
				big:      unicode.IsUpper(rune(r[0])),
			}
		}
	}

	return twelve(caves, true),twelve(caves, false)
}

func twelve(caves map[string]*node, part1 bool) int{
	var currentPath []string
	visited := map[string]struct{}{}
	count := 0
	repeatedCave := ""

	traverse(*caves["start"], visited, currentPath, &count, caves, false, &repeatedCave, part1)

	return count
}

type node struct {
	name string
	vertices []string
	big bool
}

func traverse(
	n node,
	visited map[string]struct{},
	currentPath []string,
	count *int,
	caves map[string]*node,
	doubleCaved bool,
	repeatedCave *string,
	part1 bool,
	) {
	// we can't go back to the start twice
	if  n.name == "start" && len(visited) > 0 {
		return
	}

	if _, ok := visited[n.name]; ok {
		// can't go back to a small cave in part 1
		// can't go back to more than one small cave in part 2
		if doubleCaved || part1  { return }
		doubleCaved = true
		*repeatedCave = n.name
	} else if !n.big { visited[n.name] = struct{}{} }

	currentPath = append(currentPath, n.name)

	if n.name == "end" {
		*count++
		delete(visited, n.name)
		currentPath = currentPath[:len(currentPath)-1]
		return
	}

	for _, vertex := range n.vertices {
		traverse(*caves[vertex], visited, currentPath, count, caves, doubleCaved, repeatedCave, part1)
	}

	currentPath = currentPath[:len(currentPath)-1]

	if n.name == *repeatedCave {
		*repeatedCave = ""
		doubleCaved = false
	} else {
		delete(visited, n.name)
	}
}