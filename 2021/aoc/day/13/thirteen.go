package day

import (
	"aoc/utils"
	"fmt"
	"sort"
	"strings"
)

func Thirteen(input []string) (int, int) {
	return thirteen(input, true), thirteen(input, false)
}

func thirteen(input []string, part1 bool) int {
	grid := map[int]map[int]bool{}

	for _, row := range input {
		// v lazy way of checking if it's a blank line
		if len(row) < 2 { continue }

		if row[0] == 'f' {
			// parse instruction
			row = strings.TrimPrefix(row, "fold along ")
			lineNumber := utils.ExtractInts(row)[0]
			if row[0] == 'x' {
				// iterate over grid
				for y, _ := range grid {
					for x, _ := range grid[y] {
						if x >= lineNumber {
							// add the missing values
							grid[y][lineNumber - (x - lineNumber)] = true
						}
					}
				}

				// remove the lines that we've folded over
				for y, _ := range grid {
					for x, _ := range grid[y] {
						if x >= lineNumber {
							delete(grid[y], x)
						}
					}
				}
			}

			if row[0] == 'y' {
				// iterate over grid
				for y, _ := range grid {
					if y >= lineNumber {
						currLine := lineNumber - (y - lineNumber)
						for x, _ := range grid[y] {
							// if the row doesn't exist in the grid (ie its empty)
							if _, ok := grid[currLine]; !ok {
								// add it
								grid[currLine] = map[int]bool{}
							}
							if _, ok := grid[currLine][x]; !ok {
								grid[currLine][x] = true
							}
						}
					}
				}

				for y, _ := range grid {
					if y >= lineNumber {
						delete(grid, y)
					}
				}
			}

			if part1 {
				break
			}
		} else {
			xy := utils.ExtractInts(row)
			// if the y value already existed in the grid
			if _, ok := grid[xy[1]]; ok {
				// add the x value to it
				grid[xy[1]][xy[0]] = true
			} else {
				// else create the y and add the x as its first value
				grid[xy[1]] = map[int]bool{xy[0]:true}
			}
		}
	}

	count := 0

	var keys []int
	for key, _ := range grid {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for i := 0; i <= keys[len(keys)-1]; i++ {
		if _, ok := grid[i]; ok {
			count += len(grid[i])
		}
		if !part1{
			printRow(grid[i])
		}
	}

	return count
}

func printRow(row map[int]bool) {
	var keys []int
	for key, _ := range row {
		keys = append(keys, key)
	}
	sort.Ints(keys)

	for i := 0; i < keys[len(keys)-1]; i++ {
		if _, ok := row[i]; ok {
			fmt.Printf("#")
		} else {
			fmt.Printf(".")
		}
	}
	fmt.Printf("\n")
}