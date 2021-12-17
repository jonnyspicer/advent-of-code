package utils

import (
	"regexp"
	"strconv"
)

func ExtractInts(s string) []int {
	var ints []int

	r, _ := regexp.Compile(`-?\d+`)
	matches := r.FindAllString(s, -1)

	for _, match := range matches {
		i, _ := strconv.Atoi(match)
		ints = append(ints, i)
	}

	return ints
}
