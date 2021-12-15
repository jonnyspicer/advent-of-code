package main

import (
	"aoc/day/15"
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	// will clean this up to just take an day number
	rows := parseTxt("day/15/fifteen.txt")
	a, b := day.Fifteen(rows)
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
