package main

import "fmt"

func intervalIntersection(firstList [][]int, secondList [][]int) [][]int {

	l := len(firstList)
	l1 := len(secondList)
	arr := [][]int{}

	i := 0
	j := 0

	for i < l && j < l1 {
		s, e := firstList[i][0], firstList[i][1]
		s1, e1 := secondList[j][0], secondList[j][1]
		newS := 0
		newE := 0
		if s > e1 {
			j++
		} else if s1 > e {
			i++
		} else {
			if s > s1 {
				newS = s
			} else {
				newS = s1
			}

			if e < e1 {
				newE = e
				i++
			} else {
				j++
				newE = e1
			}
			arr = append(arr, []int{newS, newE})

		}
	}

	return arr
}

func min(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func main() {

	fmt.Println(
		intervalIntersection(
			[][]int{
				{5, 6},
			},
			[][]int{
				{5, 6}, {7, 8},
			},
		),
	)
}
