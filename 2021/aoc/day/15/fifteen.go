package day

import (
	"fmt"
	"math"
	"strings"
)

func Fifteen(input []string) (int, int) {
	grid := map[coord]node{}

	dist := map[coord]int{}

	// can I use a pointer instead? or just iterate over the map?
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
					dist[c] = math.MaxInt32
				}
			}
		}
	}

	for y := 0; y < len(input) * 5; y++ {
		fmt.Printf("\n")
		for x := 0; x < len(strings.TrimSpace(input[0])) * 5; x++ {
			fmt.Printf("%v" ,grid[coord{x, y}].weight)
		}
	}

	dist[coord{0,0}] = 0

	for len(queue) > 0 {
		v := getSmallestDist(dist, visited)
		delete(queue, v)
		fmt.Println(len(queue))
		visited[v] = struct{}{}

		for _, u := range grid[v].edges {
			if dist[v] + grid[u].weight < dist[u] {
				dist[u] = dist[v] + grid[u].weight
			}
		}
	}

	return dist[coord{len(strings.TrimSpace(input[0])) * 5 - 1, len(input) *5 - 1}],0
}

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

type coord struct {
	x, y int
}

type node struct {
	edges []coord
	weight int
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