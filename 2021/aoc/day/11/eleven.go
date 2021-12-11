package day

type octopus int
type coord struct {
	x, y int
}

func Eleven(input []string) (int, int) {
	octoGrid := map[int][]octopus{}
	for y, row := range input {
		rs := []rune(row)
		for _, r := range rs {
			if r == '\r' { continue }
			octoGrid[y] = append(octoGrid[y], octopus(int(r) - '0'))
		}
	}

	flashes := 0
	firstAllFlash := 0

	for i := 0; i < 10000; i++ {
		for y, row := range octoGrid {
			for x, octo := range row {
				octoGrid[y][x] = octo.step(coord{x, y}, octoGrid)
			}
		}

		allFlash := true

		for y, row := range octoGrid {
			for x, octo := range row {
				if octo == 10 {
					flashes++
					octoGrid[y][x] = 0
				} else {
					allFlash = false
				}
			}
		}

		if allFlash {
			firstAllFlash = i
			break
		}
	}

	return flashes,firstAllFlash
}

func (o octopus) step(c coord, octoGrid map[int][]octopus) octopus {
	if o == 10 { return o }
	octoGrid[c.y][c.x] += 1
	if octoGrid[c.y][c.x] == 10 {
		ao := getAdjacentOctopi(octoGrid, c)
		for _, co := range ao {
			octoGrid[co.y][co.x] = octoGrid[co.y][co.x].step(co, octoGrid)
		}
	}

	return octoGrid[c.y][c.x]
}

func getAdjacentOctopi(m map[int][]octopus, c coord) []coord {
	adj := []coord{}

	if c.y-1 >= 0 {
		adj = append(adj, coord{c.x, c.y-1}) // top
		if c.x-1 >= 0 {
			adj = append(adj, coord{c.x-1, c.y-1}) // top left
		}
		if c.x+1 < len(m[0]) {
			adj = append(adj, coord{c.x+1, c.y-1}) // top right
		}
	}

	if c.x-1 >= 0 {
		adj = append(adj, coord{c.x-1, c.y}) // left
	}
	if c.x+1 < len(m[0]) {
		adj = append(adj, coord{c.x+1, c.y}) // right
	}

	if c.y+1 < len(m) {
		adj = append(adj, coord{c.x, c.y+1}) // bottom
		if c.x-1 >= 0 {
			adj = append(adj, coord{c.x-1, c.y+1}) // bottom left
		}
		if c.x+1 < len(m[0]) {
			adj = append(adj, coord{c.x+1, c.y+1}) // bottom right
		}
	}

	return adj
}