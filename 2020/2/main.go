package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Day Two, Part One: %v\n", dayTwo(input, a))
	fmt.Printf("Day Two, Part Two: %v\n", dayTwo(input, b))
}

func dayTwo(in []string, section func([]rune, rune, int, int) bool) int {
	valid := 0

	for _, str := range in {
		inputParts := strings.Split(str, ": ")
		policy := inputParts[0]
		password := []rune(inputParts[1])

		policyParts := strings.Split(policy, " ")
		target := []rune(policyParts[1])[0]
		bounds := strings.Split(policyParts[0], "-")
		lowerBound, _ := strconv.Atoi(bounds[0])
		upperBound, _ := strconv.Atoi(bounds[1])

		if section(password, target, lowerBound, upperBound) {
			valid++
		}
	}

	return valid
}

func a(password []rune, target rune, lowerBound int, upperBound int) bool {
	count := 0

	for _, r := range password {
		if r == target {
			count++
		}
	}

	return count >= lowerBound && count <= upperBound
}

func b(password []rune, target rune, lowerBound int, upperBound int) bool {
	return (password[lowerBound - 1] == target) != (password[upperBound - 1] == target)
}