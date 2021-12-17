package day

import (
	"aoc/utils"
	"strings"
)

type coord struct {
	x, y int
}

type probe struct {
	position coord
	xVel, yVel int
	maxHeight int
}

type targetArea struct {
	xLow, xUp, yLow, yUp int
}

func Seventeen(input []string) (int, int) {
	xy := strings.Split(strings.TrimPrefix(input[0], "target area: "), ", ")
	xs := utils.ExtractInts(xy[0])
	ys := utils.ExtractInts(xy[1])

	maxHeight := 0

	ta := targetArea{
		xLow: xs[0],
		xUp:  xs[1],
		yLow: ys[0],
		yUp:  ys[1],
	}

	availableCoords := map[coord]struct{}{}

	for x := 0; x < ta.xUp; x++ {
		// the 10000 upper bound is arbitrary and gross
		for y := ta.yLow; y < 10000; y++ {
			maxForY := 0
			p := probe {
				position: coord{0,0},
				xVel: x,
				yVel: y,
				maxHeight: 0,
			}

			b := false

			for !p.passedTargetArea(ta) {
				if p.inTargetArea(ta) {
					if _, ok := availableCoords[coord{x, y}]; !ok {
						availableCoords[coord{x, y}] = struct{}{}
					}
					if maxForY <= p.maxHeight {
						maxForY = p.maxHeight
					} else {
						b = true
					}
				}
				p.step()
				if p.xVel == 0 && p.position.x < ta.xLow { b = true; break }
			}

			if b { break }
			if maxHeight < maxForY {
				maxHeight = maxForY
			}
		}
	}

	return maxHeight,len(availableCoords)
}

func (p probe) inTargetArea(ta targetArea) bool {
	return p.position.x >= ta.xLow && p.position.x <= ta.xUp && p.position.y >= ta.yLow && p.position.y <= ta.yUp
}

func (p probe) passedTargetArea(ta targetArea) bool {
	return p.position.x > ta.xUp || p.position.y < ta.yLow
}

func (p *probe) step() {
	p.position.x += p.xVel
	p.position.y += p.yVel
	if p.xVel > 0 { p.xVel-- }
	p.yVel--
	if p.position.y > p.maxHeight {
		p.maxHeight = p.position.y
	}
}