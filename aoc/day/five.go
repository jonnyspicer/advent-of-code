package day

import (
	"strconv"
	"strings"
)

func Five(input []string) (int, int) {
	var ventLines []ventLine

	for _, vl := range input {
		ventLines = append(ventLines, stringToVentLine(vl))
	}

	return five(ventLines, false), five(ventLines, true)
}

type ventLine struct {
	x1, x2, y1, y2 int
}

func stringToVentLine(s string) ventLine {
	s = strings.TrimSuffix(s, "\r")
	coords := strings.Split(s, " -> ")
	ones := strings.Split(coords[0], ",")
	twos := strings.Split(coords[1], ",")

	x1, _ := strconv.Atoi(ones[0])
	y1, _ := strconv.Atoi(ones[1])
	x2, _ := strconv.Atoi(twos[0])
	y2, _ := strconv.Atoi(twos[1])

	return ventLine{
		x1: x1,
		x2: x2,
		y1: y1,
		y2: y2,
	}
}

func five(vls []ventLine, diagonals bool) int {
	grid := [1000][1000]int{}

	for _, v := range vls {
		v.calculateVents(&grid, diagonals)
	}

	count := 0

	for _, x := range grid {
		for _, y := range x {
			if y > 1 {count++}
		}
	}

	return count
}

func (v ventLine) calculateVents(grid *[1000][1000]int, diagonals bool) {
	if v.x1 != v.x2 && v.y1 != v.y2 {
		if diagonals {
			v.diagonal(grid)
		}
	} else {
		v.straight(grid)
	}
}

func (v ventLine) straight(grid *[1000][1000]int) {
	var xl, xu, yl, yu int

	if v.x2 < v.x1 {
		xl = v.x2
		xu = v.x1
	} else {
		xl = v.x1
		xu = v.x2
	}

	if v.y2 < v.y1 {
		yl = v.y2
		yu = v.y1
	} else {
		yl = v.y1
		yu = v.y2
	}

	for x := xl; x <= xu; x++ {
		for y := yl; y <= yu; y++ {
			grid[x][y]++
		}
	}
}

func (v ventLine) diagonal(grid *[1000][1000]int) {
		// x and y both ascend
		if v.x1 < v.x2 && v.y1 < v.y2 {
			for i := 0; i <= v.x2 - v.x1; i++ {
				grid[v.x1+i][v.y1+i]++
			}
		}

		// x ascends y descends
		if v.x1 < v.x2 && v.y1 > v.y2 {
			for i := 0; i <= v.x2 - v.x1; i++ {
				grid[v.x1+i][v.y1-i]++
			}

		}

		// x descends y ascends
		if v.x1 > v.x2 && v.y1 < v.y2 {
			for i := 0; i <= v.x1 - v.x2; i++ {
				grid[v.x1-i][v.y1+i]++
			}
		}

		// x and y both descend
		if v.x1 > v.x2 && v.y1 > v.y2 {
			for i := 0; i <= v.x1 - v.x2; i++ {
				grid[v.x1-i][v.y1-i]++
			}
		}
}