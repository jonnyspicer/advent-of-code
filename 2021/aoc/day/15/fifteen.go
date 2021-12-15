package day

import (
	"math"
	"strings"
)

type coord struct {
	x, y int
}

type node struct {
	edges []coord
	weight int
}

// This takes 45 mins to run on my machine... :(
func Fifteen(input []string) (int, int) {
	grid := map[coord]node{}

	// Pretty sure this data structure is the problem,
	// looping over it to find the smallest val is slow and gross
	dist := map[coord]int{}
	queue := map[coord]struct{}{}
	visited := map[coord]struct{}{}

	for y, row := range input {
		rs := []rune(row)
		for x, r := range rs {
			if r == '\r' { continue }

			for i := 0; i < 5; i++ {
				for j := 0; j < 5; j++ {
					c := coord{x + len(strings.TrimSpace(input[0]))*i, y + len(input)*j}

					weight := int(r) - '0' + i + j
					if weight > 9 {
						weight = 1 + weight % 10
					}

					grid[c] = newNode(c, coord{len(strings.TrimSpace(input[0])) * 5, len(input) * 5}, weight)
					queue[c] = struct{}{}
				}
			}
		}
	}

	// set initial node to top left
	dist[coord{0,0}] = 0

	for len(queue) > 0 {
		v := getSmallestDist(dist, visited)
		delete(queue, v)
		visited[v] = struct{}{}

		for _, u := range grid[v].edges {
			if _, ok := dist[u]; !ok || dist[v] + grid[u].weight < dist[u] {
				dist[u] = dist[v] + grid[u].weight
			}
		}
	}

	return dist[coord{len(strings.TrimSpace(input[0])) - 1, len(input) - 1}],dist[coord{len(strings.TrimSpace(input[0])) * 5 - 1, len(input) *5 - 1}]
}

// this is the problem... takes ~11ms per call when dist gets large
func getSmallestDist(dist map[coord]int, visited map[coord]struct{}) coord {
	smallestDist := math.MaxInt32
	var smallestCoord coord

	for key, val := range dist {
		if _, ok := visited[key]; !ok {
			if val < smallestDist {
				smallestDist = val
				smallestCoord = key
			}
		}
	}

	return smallestCoord
}

func newNode(c, max coord, weight int) node {
	edges := []coord{}

	if c.x > 0 {
		edges = append(edges, coord{c.x-1, c.y})
	}

	if c.x < max.x {
		edges = append(edges, coord{c.x+1, c.y})
	}

	if c.y > 0 {
		edges = append(edges, coord{c.x, c.y-1})
	}

	if c.y < max.y {
		edges = append(edges, coord{c.x, c.y+1})
	}

	return node{edges, weight}
}