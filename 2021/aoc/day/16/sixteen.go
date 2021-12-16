package day

import (
	"fmt"
	"strconv"
)

var comparisons = map[int]func(int, int)int{
	0: sum,
	1: prod,
	2: min,
	3: max,
	5: greater,
	6: lesser,
	7: equal,
}

func Sixteen(input []string) (int, int) {
	m := map[rune][]rune {
		'0': {'0','0','0','0'},
		'1': {'0','0','0','1'},
		'2': {'0','0','1','0'},
		'3': {'0','0','1','1'},
		'4': {'0','1','0','0'},
		'5': {'0','1','0','1'},
		'6': {'0','1','1','0'},
		'7': {'0','1','1','1'},
		'8': {'1','0','0','0'},
		'9': {'1','0','0','1'},
		'A': {'1','0','1','0'},
		'B': {'1','0','1','1'},
		'C': {'1','1','0','0'},
		'D': {'1','1','0','1'},
		'E': {'1','1','1','0'},
		'F': {'1','1','1','1'},
	}
	hex := []rune(input[0])
	binarunes := []rune{}
	for _, r := range hex {
		binarunes = append(binarunes, m[r]...)
	}

	i, verSum, valSum := 0, 0, -1

	// while we're not at the end
	for i < len(binarunes) {
		subVerSum, subValSum, endPos := parsePacket(binarunes, i)

		verSum += subVerSum
		if subValSum >= 0 {
			valSum = subValSum
		}
		i = endPos
	}

	return verSum,valSum
}

// returns the sum of the versionTypes of this packet and any subpackets, its value and the end position
// i is our starting position
func parsePacket(binarunes []rune, i int) (int, int, int) {
	versionSum, valueSum := 0, -1

	// slightly gross way to ignore trailing zeroes
	// 11 is shortest possible packet
	if i + 11 > len(binarunes) {
		return versionSum, valueSum, len(binarunes)
	}

	versionType := br(binarunes[i:i+3])
	typeID := br(binarunes[i+3:i+6])

	i += 6
	versionSum += versionType

	if typeID == 4 {
		value, endPos := packetLiteral(binarunes, i)
		valueSum += value + 1
		i = endPos
	} else {
		subVerSum, subValSum, endPos := operator(binarunes, i, typeID)
		valueSum = comparisons[typeID](valueSum, subValSum)
		i = endPos
		versionSum += subVerSum
	}

	return versionSum, valueSum, i
}

// turns an array of bit runes into an int
func br(r []rune) int {
	i, err := strconv.ParseInt(string(r), 2, 64)
	if err != nil {
		fmt.Printf("C'est cassé: %v", err)
	}
	return int(i)
}

// returns the value of the packet and its last position
func packetLiteral(binarunes []rune, i int) (int, int) {
	var rs []rune
	for i < len(binarunes) {
		rs = append(rs, binarunes[i+1:i+5]...)
		if binarunes[i] == '0' {
			return br(rs), i + 5
		}
		i += 5
	}

	return 0,0
}

// returns the sum of the versionTypes and values of this packet and any subpackets and the end position
// i is our starting position
func operator(binarunes []rune, i, typeID int) (int, int, int) {
	versionSum, value := 0, -1
	if binarunes[i] == '0' {
		// next 15 bits are a number that represents
		// the total length in bits of the sub-packets contained
		// by the packet
		subPacketBitLength := br(binarunes[i+1:i+16])
		i += 16
		j := 0
		for j < subPacketBitLength {
			subVerSum, subVal, endPos := parsePacket(binarunes, i)
			versionSum += subVerSum
			value = comparisons[typeID](value, subVal)
			j += endPos - i
			i = endPos
		}
	} else {
		// next 11 bits are a number that represents the number
		// of sub-packets immediately contained by this packet
		numberOfSubPackets := br(binarunes[i+1:i+12])
		i += 12
		j := 0
		for j < numberOfSubPackets {
			subVerSum, subVal, endPos := parsePacket(binarunes, i)
			versionSum += subVerSum
			value = comparisons[typeID](value, subVal)
			i = endPos
			j++
		}
	}

	return versionSum, value, i
}

func sum(a int, b int) int {
	if a < 0 { a = 0 }
	return a + b
}

func min(a int, b int) int {
	if a < 0 { return b }
	if a < b { return a }
	return b
}

func max(a int, b int) int {
	if a < 0 { return b }
	if a > b { return a }
	return b
}

func prod(a int, b int) int {
	if a < 0 { return b }
	return a * b
}

func greater(a int, b int) int {
	if a < 0 { return b }
	if a > b { return 1 }
	return 0
}

func lesser(a int, b int) int {
	if a < 0 { return b }
	if a < b { return 1 }
	return 0
}

func equal(a int, b int) int {
	if a < 0 { return b }
	if a == b { return 1 }
	return 0
}

