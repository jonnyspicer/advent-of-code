package day

import (
	"aoc/utils"
	"math"
	"sort"
)

func Seven(input []string) (int, int) {
	positions := utils.ExtractInts(input[0])
	med := median(positions)
	mea := mean(positions)
	medFuel, meaFuel := 0, 0

	for _, p := range positions {
		medFuel += int(math.Abs(float64(p - med)))
		for i := 0; i <= int(math.Abs(float64(p - mea))); i++ {
			meaFuel += i
		}
	}

	return medFuel, meaFuel
}

func median(input []int) int {
	sort.Ints(input)

	middle := len(input) / 2

	if middle % 2 == 1 {
		return input[middle]
	}

	return (input[middle-1] + input[middle]) / 2
}

func mean(input []int) int {
	sum := 0

	for _, i := range input {
		sum += i
	}

	return sum / len(input)
}