package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	m := make(map[int]int)
	arr := [][]int{}
	sort.Ints(nums)
	for i, v := range nums {
		m[v] = i
	}

	ci := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, ok := ci[nums[i]]; !ok {
			ci[nums[i]] = true

			cj := make(map[int]bool)

			for j := i + 1; j < len(nums); j++ {
				if _, ok := cj[nums[j]]; !ok {
					cj[nums[j]] = true

					sum := nums[i] + nums[j]
					if v, ok := m[-sum]; ok && v > j && v != i && v != j {
						arr = append(arr, []int{nums[i], nums[j], -sum})
					}
				}
			}
		}
	}

	return arr

}

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}
