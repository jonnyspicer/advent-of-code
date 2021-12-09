package day

import (
	"sort"
)

type coord struct {
	x int
	y int
}

func Nine(input []string) (int, int) {
	grid := map[int][]int{}

	for i, line := range input {
		// convert to runes and save in slice
		rs := []rune(line)
		for _, r := range rs {
			if r == '\r' { continue }
			grid[i] = append(grid[i], int(r) - '0')
		}
	}

	sum := 0

	var lowPoints []coord

	// down the grid
	for i := 0; i < len(grid); i++ {
		// across the grid
		for j := 0; j < len(grid[0]); j++ {
			if lowPoint(grid, coord{j, i}) {
				lowPoints = append(lowPoints, coord{j, i})
				sum += 1 + grid[i][j]
			}
		}
	}

	var sizes []int

	// while there are points to check
	for _, lp := range lowPoints {
		// mark it as seen
		// check the other points, add them to seen
		// add them to list to be checked as well

		var basin []coord
		toBeChecked := []coord{lp}

		for len(toBeChecked) > 0 {
			for i := len(toBeChecked) - 1; i >= 0; i-- {
				tbc := toBeChecked[i]
				// check if its part of the basin
				// if it is, add it to the basin and add the relevant elements to tbc
				// remove tbc from to be checked

				if !contains(tbc, basin) && isInBounds(grid, tbc) && grid[tbc.y][tbc.x] != 9 {
					basin = append(basin, tbc)
					up := coord{tbc.x, tbc.y-1}
					down := coord{tbc.x, tbc.y+1}
					left := coord{tbc.x-1, tbc.y}
					right := coord{tbc.x+1, tbc.y}
					s := []coord{up, down, left, right}
					toBeChecked = append(toBeChecked, s...)
				}

				toBeChecked[i] = toBeChecked[len(toBeChecked)-1]
				toBeChecked = toBeChecked[:len(toBeChecked)-1]
			}
		}

		sizes = append(sizes, len(basin))
	}

	sort.Ints(sizes)

	return sum, sizes[len(sizes) - 1] * sizes[len(sizes) - 2] * sizes[len(sizes) - 3]
}

func isInBounds(m map[int][]int, c coord) bool {
	return c.x >= 0 && c.y >= 0 && c.x < len(m[0]) && c.y < len(m)
}

func lowPoint(m map[int][]int, c coord) bool {
	i, j := c.y, c.x
	if !isInBounds(m, coord{j, i}) { return false }

	var up, left, right, down bool
	// up
	if i > 0 {
		up = m[i][j] < m[i-1][j]
	} else {
		up = true
	}

	// left
	if j > 0 {
		left = m[i][j] < m[i][j-1]
	} else {
		left = true
	}

	// right
	if j < len(m[i]) - 1 {
		right = m[i][j] < m[i][j+1]
	} else {
		right = true
	}

	// down
	if i < len(m) - 1 {
		down = m[i][j] < m[i+1][j]
	} else {
		down = true
	}

	return up && left && right && down
}

func contains(needle coord, haystack []coord) bool {
	for _, hay := range haystack {
		if hay.x == needle.x && hay.y == needle.y { return true }
	}
	return false
}