package day

import (
	"fmt"
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

	g, err := strconv.ParseInt(string(gamma), 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	e, err := strconv.ParseInt(string(epsilon), 2, 64)
	if err != nil {
		fmt.Println(err)
	}

	// part two

	// maps to store which the indexes of the numbers we're still interested in
	oxygenNumbers := map[int]bool{}
	co2Numbers := map[int]bool{}
	for k := 0; k < len(input); k++ {
		oxygenNumbers[k] = true
		co2Numbers[k] = true
	}

	// for every bit, starting from left
	for l := 0; l < length; l++ {
		count := 0
		mcb := mostCommonBit(oxygenNumbers, input, l)
		// for every binary number in the input
		for n, in := range input {
			// if it's currently true and the bit is either the most common OR a 1 when there is no mcb
			if (
				(rune(in[l]) == '1' && mcb > 0) ||
					(rune(in[l]) == '0' && mcb < 0) ||
					(rune(in[l]) == '1' && mcb == 0)) &&
				oxygenNumbers[n] == true {
				oxygenNumbers[n] = true
			} else {
				// we're no longer interested in this number
				oxygenNumbers[n] = false
			}
		}

		// loop over our map again to see if we only have one number left
		for _, oxyNum := range oxygenNumbers {
			if oxyNum == true {
				count++
			}
		}

		if count == 1 {
			break
		}
	}

	for l := 0; l < length; l++ {
		count := 0
		mcb := mostCommonBit(co2Numbers, input, l)
		for n, in := range input {
			if (
				(rune(in[l]) == '1' && mcb < 0) ||
					(rune(in[l]) == '0' && mcb > 0) ||
					(rune(in[l]) == '0' && mcb == 0)) &&
				co2Numbers[n] == true {
				co2Numbers[n] = true
			} else {
				co2Numbers[n] = false
			}
		}
		for _, co2Num := range co2Numbers {
			if co2Num == true {
				count++
			}
		}

		if count == 1 {
			break
		}
	}

	var gstr, estr string

	for orating, oval := range oxygenNumbers {
		if oval {
			gstr = input[orating]
		}
	}

	for crating, cval := range co2Numbers {
		if cval {
			estr = input[crating]
		}
	}

	g2, err := strconv.ParseInt(gstr[0:length], 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	e2, err := strconv.ParseInt(estr[0:length], 2, 64)
	if err != nil {
		fmt.Println(err)
	}

	return g * e, g2 * e2
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