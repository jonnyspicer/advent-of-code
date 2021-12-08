package day

import (
	"math"
	"strings"
)

var requiredSegments = map[int][]int{
	0: {1,2,3,5,6,7},
	1: {3,6},
	2: {1,3,4,5,7},
	3: {1,3,4,6,7},
	4: {2,3,4,6},
	5: {1,2,4,6,7},
	6: {1,2,4,5,6,7},
	7: {1,3,6},
	8: {1,2,3,4,5,6,7},
	9: {1,2,3,4,6,7},
}

func Eight(input []string) (int, int) {
	easyCount, hardCount := 0, 0
	for _, in := range input {
		parts := strings.Split(in, " | ")
		wires := strings.Fields(parts[0])
		output := strings.Fields(parts[1])

		// all possible values around the digit face
		possibleValues := map[int][]string{
			// top
			1: {"a", "b", "c", "d", "e", "f", "g"},
			// top left
			2: {"a", "b", "c", "d", "e", "f", "g"},
			// top right
			3: {"a", "b", "c", "d", "e", "f", "g"},
			// middle
			4: {"a", "b", "c", "d", "e", "f", "g"},
			// bottom left
			5: {"a", "b", "c", "d", "e", "f", "g"},
			// bottom right
			6: {"a", "b", "c", "d", "e", "f", "g"},
			// bottom
			7: {"a", "b", "c", "d", "e", "f", "g"},
		}

		dw := [10]bool{}
		// while we haven't figured out each of the inputs
		for decodedWires(dw) < len(wires) {
			// loop through them
			for i, w := range wires {
				if dw[i] == true {
					continue
				}

				switch len(w) {
				case 2:
					// digit is a 1
					possibleValues = updatePossibleValues(1, w, possibleValues)
					dw[i] = true

				case 3:
					// digit is a 7
					possibleValues = updatePossibleValues(7, w, possibleValues)
					dw[i] = true

				case 4:
					// digit is a 4
					possibleValues = updatePossibleValues(4, w, possibleValues)
					dw[i] = true

				case 5:
					// digit could be a 2, 3, 5
					two := couldBe(possibleValues, w, requiredSegments[2])
					three := couldBe(possibleValues, w, requiredSegments[3])
					five := couldBe(possibleValues, w, requiredSegments[5])

					if two && !three && !five {
						// it's a two
						possibleValues = updatePossibleValues(2, w, possibleValues)
						dw[i] = true
					} else if !two && three && !five {
						// it's a three
						possibleValues = updatePossibleValues(3, w, possibleValues)
						dw[i] = true
					} else if !two && !three && five {
						// it's a five
						possibleValues = updatePossibleValues(5, w, possibleValues)
						dw[i] = true
					}

				case 6:
					// digit could be a 0, 6 or 9
					zero := couldBe(possibleValues, w, requiredSegments[0])
					six := couldBe(possibleValues, w, requiredSegments[6])
					nine := couldBe(possibleValues, w, requiredSegments[9])

					if zero && !six && !nine {
						// it's a zero
						possibleValues = updatePossibleValues(0, w, possibleValues)
						dw[i] = true
					} else if !zero && six && !nine {
						// it's a six
						possibleValues = updatePossibleValues(6, w, possibleValues)
						dw[i] = true
					} else if !zero && !six && nine {
						// it's a nine
						possibleValues = updatePossibleValues(9, w, possibleValues)
						dw[i] = true
					}

				case 7:
					// digit is an 8
					dw[i] = true
				}
			}
		}

		sum := 0
		for i, o := range output {
			switch len(o) {
			case 2, 4, 3, 7:
				easyCount++
			}

			// raises the digit by a power of 10 to put it in the right place
			// rather than faff around with type conversions and concatenation
			sum += decodeDigit(o, possibleValues) * int(math.Pow(10, float64(len(output) - i - 1)))
		}

		hardCount += sum
	}
	return easyCount,hardCount
}

func updatePossibleValues(i int, s string, possibleValues map[int][]string) map[int][]string{
	// for every required segment
	for j := 0; j < len(requiredSegments[i]); j++ {
		// remove all the values that are no longer possible
		possibleValues[requiredSegments[i][j]] = reduceSlice(possibleValues[requiredSegments[i][j]], s)
	}

	return possibleValues
}

func decodedWires(s [10]bool) int {
	count := 0
	for _, b := range s {
		if b { count ++ }
	}
	return count
}

func decodeDigit(s string, m map[int][]string) int {
	switch len(s) {
	case 2:
		return 1
	case 3:
		return 7
	case 4:
		return 4
	case 5:
		if couldBe(m, s, requiredSegments[2]) { return 2 }
		if couldBe(m, s, requiredSegments[3]) { return 3 }
		if couldBe(m, s, requiredSegments[5]) { return 5 }
	case 6:
		if couldBe(m, s, requiredSegments[0]) { return 0 }
		if couldBe(m, s, requiredSegments[6]) { return 6 }
		if couldBe(m, s, requiredSegments[9]) { return 9 }
	}
	return 8
}

// removes all the chars from the slice that aren't present in the matching string
// doesn't necessarily reduce down to a single char like a javascript array.reduce()
func reduceSlice(sl []string, st string) []string {
	// for every element in the slice
	for i := len(sl)-1; i >= 0; i-- {
		// if the wire doesn't contain that element
		if !strings.Contains(st, sl[i]) {
			sl = removeElement(sl, i)
		}
	}

	return sl
}

func removeElement(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// takes a map of all the current possible values, the wire string,
// the number we're interested in seeing if it could be and a slice of required values for that number
func couldBe(m map[int][]string, w string, s []int) bool {
	seen := map[int]bool{}
	reps := len(w)

	// the number of iterations here is fairly arbitrary
	for j := 0; j < reps; j++ {
		// for each value we need to fill
		for _, val := range s {
			// see how many options we have
			mr := matchingRunes([]rune(w), m[val])
			switch len(mr) {
			// if there's only one, remove that char from the string before continuing
			case 1:
				w = strings.Replace(w, string(mr[0]), "", -1)
				seen[val] = true
			// if there's none
			case 0:
				// and we haven't already filled the value
				if !seen[val] {
					return false
				}
			// if we still have multiple options left for multiple values
			default:
				// And aren't going to make any more progress by iterating more
				if j == reps - 1 {
					return true
				}
			}
		}
	}

	// if we don't have all the values we need
	for _, val := range s {
		if !seen[val] { return false }
	}

	return true
}

// gets all runes that appear in both slices
func matchingRunes(needles []rune, haystack []string) []rune {
	var rs []rune
	// for every character in the wire string
	for _, r := range needles {
		// for every element in the relevant map
		for _, elem := range haystack {
			if string(r) == elem { rs = append(rs, r) }
		}
	}
	return rs
}