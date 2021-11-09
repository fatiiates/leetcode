package main

import (
	"fmt"
)

func findNumOfValidWords(words []string, puzzles []string) []int {
	masked := make(map[int]int)

	for _, w := range words {

		b := mask(w)
		if _, ok := masked[b]; !ok {
			masked[b] = 1
		} else {
			masked[b] += 1
		}
	}

	result := make([]int, len(puzzles))
	for i, p := range puzzles {

		m := mask(p)

		first := 1 << (p[0] - 97)

		for j := m; j != 0; j = ((j - 1) & m) {

			if (j & first) == 0 {
				continue
			}
			if v, ok := masked[j]; ok {
				result[i] += v
			}
		}

	}
	return result
}

func mask(s string) int {
	b := 0
	for _, l := range s {
		b |= 1 << (l - 97)
	}
	return b
}

func main() {

	fmt.Println(findNumOfValidWords([]string{"aaaa", "asas"}, []string{"absoveyz"}))
}
