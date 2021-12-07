package day

func One(readings []int, gap int) int {
	count := 0

	for i := gap; i < len(readings); i++ {
		if readings[i] > readings[i - gap] {
			count++
		}
	}

	return count
}