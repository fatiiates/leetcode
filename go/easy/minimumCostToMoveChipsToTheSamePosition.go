package main

func minCostToMoveChips(position []int) int {

	odds := 0
	evens := 0
	l := len(position)
	for i := 0; i < l; i++ {
		if position[i]&1 == 1 {
			odds++
		} else {
			evens++
		}
	}

	return min(odds, evens)
}

func min(n1 int, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}
