package main

import "fmt"

func main() {
	fmt.Printf("Day Two, Part One: %v\n", dayOneA(input))
	fmt.Printf("Day Two, Part Two: %v\n", dayOneB(input))
}

func dayOneA(input []int) int {
	m := map[int]struct{}{}

	for _, k := range input {
		if _, ok := m[k]; ok {
			return k * (2020 - k)
		}
		m[2020 - k] = struct{}{}
	}

	return 0
}

func dayOneB(input []int) int {
	m := map[int][]int{}

	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			if i == j {
				continue
			}
			if input[i] + input[j] < 2020 {
				m[2020 - input[i] - input[j]] = []int{input[i], input[j]}
			}
		}
	}

	for _, k := range input {
		if _, ok := m[k]; ok {
			return k * m[k][0] * m[k][1]
		}
	}

	return 0
}