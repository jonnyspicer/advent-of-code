package main

import (
	"aoc/day"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	// will clean this up to just take an day number
	rows := parseTxt("day/one.txt")
	fmt.Printf("Part one: %v\n", day.One(stringToIntSlice(rows), 1))
	fmt.Printf("Part two: %v\n", day.One(stringToIntSlice(rows), 3))
}

func parseTxt(path string) []string {
	rows, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}

	return strings.Split(string(rows), "\n")
}

func stringToIntSlice(strings []string) []int {
	var ints []int

	for _, str := range strings {
		i, _ := strconv.Atoi(str)
		ints = append(ints, i)
	}

	return ints
}