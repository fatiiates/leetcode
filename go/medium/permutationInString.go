package main

import (
	"fmt"
	"reflect"
	"sort"
)

func checkInclusion(s1 string, s2 string) bool {

	l1 := len(s1)
	l2 := len(s2)
	if l2 < l1 {
		return false
	}

	var s1Map map[byte]int16 = map[byte]int16{}
	for i := 0; i < l1; i++ {
		if _, ok := s1Map[s1[i]]; ok {
			s1Map[s1[i]] += 1
		} else {
			s1Map[s1[i]] = 1
		}
	}

	sortedS1 := []byte(s1)
	sort.Slice(sortedS1, func(i, j int) bool {
		return sortedS1[i] < sortedS1[j]
	})

	for i := 0; i <= l2-l1; i++ {

		if compare(&s1Map, s2[i]) {
			s := s2[i : i+l1]

			sortedS := []byte(s)
			sort.Slice(sortedS, func(i, j int) bool {
				return sortedS[i] < sortedS[j]
			})

			if reflect.DeepEqual(sortedS, sortedS1) {
				return true
			}
		}

	}
	return false

}

func compare(m *map[byte]int16, l byte) bool {
	if _, ok := (*m)[l]; ok {
		return true
	}
	return false
}

func main() {

	fmt.Println(checkInclusion("acb", "eidbaooo"))

}
