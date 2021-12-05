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
	rows := parseTxt("day/five.txt")
	a, b := day.Five(rows)
	fmt.Printf("Part one: %v\n", a)
	fmt.Printf("Part two: %v\n", b)
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