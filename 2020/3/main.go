package main

import "fmt"

func main() {
	fmt.Printf("Day Three, Part One: %v\n", dayThree(input, 3, 1))
	fmt.Printf("Day Three, Part One: %v\n",
		dayThree(input, 1, 1) *
			dayThree(input, 3, 1) *
			dayThree(input, 5, 1) *
			dayThree(input, 7, 1) *
			dayThree(input, 1, 2),
			)
}

func dayThree(in []string, right int, down int) int {
	trees := 0
	rowLength := len(in[0])

	for i := 0; i < len(in); i += down {
		row := []rune(in[i])
		if row[((i / down) * right) % rowLength] == '#' {
			trees++
		}
	}

	return trees
}