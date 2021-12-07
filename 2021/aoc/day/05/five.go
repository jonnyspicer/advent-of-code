package day

import (
	"aoc/utils"
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
	ints := utils.ExtractInts(s)

	return ventLine{
		x1: ints[0],
		x2: ints[2],
		y1: ints[1],
		y2: ints[3],
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
			if y > 1 { count++ }
		}
	}

	return count
}

func (v ventLine) calculateVents(grid *[1000][1000]int, diagonals bool) {
	if v.x1 != v.x2 && v.y1 != v.y2 {
		if diagonals { v.diagonal(grid) }
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