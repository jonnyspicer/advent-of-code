package day

import (
	"strconv"
)

func Three(input []string) (int64, int64) {
	length := 12

	// part 1

	gamma := make([]rune, length)
	epsilon := make([]rune, length)
	m := make(map[int]int, length)

	for _, row := range input {
		bits := []rune(row)

		for i := 0; i < length; i++ {
			if bits[i] == '1' {
				m[i]++
			} else {
				m[i]--
			}
		}
	}

	for j := 0; j < length; j++ {
		if m[j] > 0 {
			gamma[j] = '1'
			epsilon[j] = '0'
		} else {
			gamma[j] = '0'
			epsilon[j] = '1'
		}
	}

	g, _ := strconv.ParseInt(string(gamma), 2, 64)
	e, _ := strconv.ParseInt(string(epsilon), 2, 64)

	// part two

	oxygenNumbers := lastRemainingString(input, length, oxygen)
	co2Numbers := lastRemainingString(input, length, carbon)

	oxygenString := ratingFromMap(oxygenNumbers, input)
	carbonString := ratingFromMap(co2Numbers, input)

	o, _ := strconv.ParseInt(oxygenString[0:length], 2, 64)
	c, _ := strconv.ParseInt(carbonString[0:length], 2, 64)

	return g * e, o * c
}
// returns 0 if there is no mcb, a positive value if the mcb is 1 and a negative value if the mcb is 0
func mostCommonBit(m map[int]bool, input []string, position int) int {
	count := 0

	for key, val := range m {
		if val {
			if input[key][position] == '1' {
				count++
			} else {
				count--
			}
		}
	}

	return count
}

func oxygen(r rune, mcb int) bool {
	return (r == '1' && mcb > 0) || (r == '0' && mcb < 0 ) || r == '1' && mcb == 0
}

func carbon(r rune, mcb int) bool {
	return (r == '1' && mcb < 0) || (r == '0' && mcb > 0 ) || r == '0' && mcb == 0
}

// will return a map with only a single value of true, where the corresponding key is the index
// in the input slice that we're interested in
func lastRemainingString(input []string, length int, eval func(r rune, mcb int) bool) map[int]bool {
	// first create a map of all the indexes where every value is true
	// seeing as we want to start with all the binary numbers
	m := map[int]bool{}
	for k := 0; k < len(input); k++ {
		m[k] = true
	}

	for l := 0; l < length; l++ {
		count := 0
		mcb := mostCommonBit(m, input, l)
		// for every binary number in the input
		for n, in := range input {
			// if our evaluation is false or if the value in the map is already false
			if !eval(rune(in[l]), mcb) || !m[n] {
				m[n] = false
			}
		}

		// loop over our map again to see if we only have one number left
		for _, val := range m {
			if val {
				count++
			}
		}

		if count == 1 {
			break
		}
	}

	return m
}

func ratingFromMap(m map[int]bool, input []string) string {
	for k, v := range m {
		if v {
			return input[k]
		}
	}

	return ""
}