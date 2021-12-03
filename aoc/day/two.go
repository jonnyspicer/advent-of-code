package day

import (
	"strconv"
	"strings"
)

func Two(input []string, navigate func(x, y, z *int, direction string, distance int)) int {
	x, y, z := 0, 0, 0

	for _, move := range input {
		directions := strings.Split(move, " ")
		distance, _ := strconv.Atoi(directions[1])
		navigate(&x, &y, &z, directions[0], distance)
	}

	return x * y
}

func A(x, y, z *int, direction string, distance int) {
	switch direction {
	case "forward":
		*x += distance
	case "down":
		*y += distance
	case "up":
		*y -= distance
	}
}

func B(x, y, z *int, direction string, distance int) {
	switch direction {
	case "forward":
		*x += distance
		*y += distance * *z
	case "down":
		*z += distance
	case "up":
		*z -= distance
	}
}
