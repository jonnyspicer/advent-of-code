package day

import (
	"fmt"
	"sort"
	"strings"
)

func Fourteen(input []string) (int, int) {
	return fourteen(input, 10), fourteen(input, 40)
}

func fourteen(input []string, iterations int) int {
	polymer := strings.TrimSpace(input[0])
	mappings := map[string][]string{}
	pairs := map[string]int{}
	letters := map[rune]int{}

	for i := 2; i < len(input); i++ {
		mapping := strings.Split(input[i], " -> ")
		// for each pair, store the two other pairs it will create after the insertion
		mappings[mapping[0]] = []string{
			mapping[0][:1] + strings.TrimSpace(mapping[1]),
			strings.TrimSpace(mapping[1]) + mapping[0][1:],
		}

		// Add a new entry in the pairs map
		pairs[mapping[0]] = 0
	}

	// store the letters from the initial polymer string
	for _, char := range polymer {
		if _, ok := letters[char]; !ok {
			letters[char] = 1
		} else {
			letters[char]++
		}
	}

	// store the initial pairs
	for j := 0; j < len(polymer) - 1; j++ {
		key := string(polymer[j]) + string(polymer[j+1])
		pairs[key]++
	}

	for i := 0; i < iterations; i++ {
		temp := map[string]int{}

		// for every pair
		for pair, pairCount := range pairs {
			if pairCount == 0 { continue }

			// the pairs we'll need to update
			toUpdate := mappings[pair]

			// add the appropriate number of newly created letters
			letters[rune(toUpdate[0][1])] += pairCount

			// store number of extra pairs created in temp map
			for key := range toUpdate {
				if _, ok := temp[toUpdate[key]]; !ok {
					temp[toUpdate[key]] = pairCount
				} else {
					temp[toUpdate[key]] += pairCount
				}
			}
		}

		// copy the values from the temp map into the main one
		for key := range pairs {
			if _, ok := temp[key]; ok {
				pairs[key] = temp[key]
			} else {
				// if there is no value in the temp map, set the main one to 0
				pairs[key] = 0
			}
		}
	}

	var vals []int

	for _, val := range letters {
		vals = append(vals, val)
	}

	sort.Ints(vals)

	return vals[len(vals)-1] - vals[0]
}

func naive(input []string) int {
	polymer := strings.TrimSpace(input[0])

	rules := map[string]string{}

	for i := 2; i < len(input); i++ {
		mapping := strings.Split(input[i], " -> ")
		rules[mapping[0]] = strings.TrimSpace(mapping[1])
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		for j := len(polymer)-1; j > 0; j-- {
			key := string(polymer[j-1]) + string(polymer[j])

			if substr, ok := rules[key]; ok {
				polymer = polymer[:j] + substr + polymer[j:]
			}
		}
	}

	letters := map[rune]int {}

	for _, char := range polymer {
		if _, ok := letters[char]; !ok {
			letters[char] = 1
		} else {
			letters[char]++
		}
	}

	var vals []int

	for _, val := range letters {
		vals = append(vals, val)
	}

	sort.Ints(vals)

	return vals[len(vals)-1] - vals[0]
}