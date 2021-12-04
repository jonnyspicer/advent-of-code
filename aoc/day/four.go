package day

import (
	"encoding/json"
	"strconv"
	"strings"
)

type bingo struct {
	values [25]int
	called [25]bool
	turnsToWin int
	lastNumberCalled int
	unmarkedSum int
}

func Four(input []string) (int, int) {
	bingos := []bingo{}

	m := input[0]

	for i := 2; i < len(input); i += 6 {
		bingoStringo := strings.Join(input[i:i+5], "")
		bingos = append(bingos, stringoToBingo(bingoStringo, m))
	}

	fastestBingo := bingos[0]
	slowestBingo := bingos[0]

	for _, b := range bingos {
		if b.turnsToWin < fastestBingo.turnsToWin {
			fastestBingo = b
		}
		if b.turnsToWin > slowestBingo.turnsToWin {
			slowestBingo = b
		}
	}

	return fastestBingo.lastNumberCalled * fastestBingo.unmarkedSum, slowestBingo.lastNumberCalled * slowestBingo.unmarkedSum
}

func (b bingo) isWinner() bool {
	for i := 0; i < 5; i++ {
		if (b.called[5 * i] && b.called[5*i+1] && b.called[5*i+2] && b.called[5*i+3] && b.called[5*i+4]) ||
			(b.called[i] && b.called[i + 5] && b.called[i + 10] && b.called[i + 15] && b.called[i + 20]) {
			return true
		}
	}

	return false
}

func stringoToBingo(s, m string) bingo {
	b := bingo{}
	vals := strings.Fields(s)
	for i := 0; i < 25; i++ {
		b.values[i], _ = strconv.Atoi(vals[i])
	}
	b.numberOfTurnsToWin(m)
	return b
}

func (b *bingo) numberOfTurnsToWin(nums string) {
	// parse nums
	// for each one, mark as seen && check winner
	// if winner, mark last number
	// calculate everything else

	fakeJson := "[" + nums + "]"
	var moves []int
	_ = json.Unmarshal([]byte(fakeJson), &moves)

	// for every move
	for i := 0; i < len(moves); i++ {
		// for every value
		for j := 0; j < len(b.values); j++ {
			if moves[i] == b.values[j] {
				b.called[j] = true
				if b.isWinner() {
					b.lastNumberCalled = moves[i]
					b.turnsToWin = i
					b.sum()
					return
				}
			}
		}
	}
}

func (b *bingo) sum() {
	sum := 0

	for i := 0; i < 25; i++ {
		if !b.called[i] {
			sum += b.values[i]
		}
	}

	b.unmarkedSum = sum
}