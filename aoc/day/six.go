package day

import (
	"aoc/utils"
	"fmt"
)

func Six(input []string) (int, int) {
	days := utils.ExtractInts(input[0])

	return fishPop(days, 80), fishPop(days, 256)
}

func fishPop(in []int, days int) int {
	lf := map[int]int{
		0: 0,
		1: 0,
		2: 0,
		3: 0,
		4: 0,
		5: 0,
		6: 0,
		7: 0,
		8: 0,
	}

	for _, timer := range in {
		lf[timer]++
	}

	for i := 0; i < days; i++ {
		fmt.Println(i)
		t0, t1, t2, t3, t4, t5, t6, t7, t8 := lf[0], lf[1], lf[2], lf[3], lf[4], lf[5], lf[6], lf[7], lf[8]
		lf[8] = t0
		lf[7] = t8
		lf[6] = t7 + t0
		lf[5] = t6
		lf[4] = t5
		lf[3] = t4
		lf[2] = t3
		lf[1] = t2
		lf[0] = t1
	}

	return lf[0] + lf[1] + lf[2] + lf[3] + lf[4] + lf[5] + lf[6] + lf[7] + lf[8]
}

// my first try, included for posterity
func SlowSix(input []string) (int, int) {
	var lf []lanternfish
	for _, days := range utils.ExtractInts(input[0]) {
		lf = append(lf, newFish(days))
	}

	for i := 0; i < 80; i++ {
		for _, fish := range lf {
			fish.processDay(&lf)
		}
	}

	return len(lf), 0
}

type lanternfish struct {
	daysUntilSpawn int
}

func newFish(dus int) lanternfish {
	return lanternfish{ daysUntilSpawn: dus }
}

func (lf *lanternfish) processDay(fishList *[]lanternfish) {
	*lf = newFish(lf.daysUntilSpawn - 1)
	if lf.daysUntilSpawn < 0 {
		lf.daysUntilSpawn = 6
		*fishList = append(*fishList, newFish(8))
	}
}