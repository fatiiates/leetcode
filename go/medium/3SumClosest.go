package main

import (
	"fmt"
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	m := make(map[int]int)
	closest := math.MaxInt32
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
					for k := j + 1; k < len(nums); k++ {
						tsum := sum + nums[k]
						closest = min(tsum, closest, target)
						if k < len(nums)-1 && abs(target-tsum) <= abs(target-(sum+nums[k+1])) {
							break
						}
					}
				}
			}
		}
	}

	return closest

}

func min(m, n, t int) int {
	if abs(t-m) < abs(t-n) {
		return m
	}
	return n
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func main() {
	fmt.Println(threeSumClosest([]int{1, 2, 4, 8, 16, 32, 64, 128}, 82))
}

/*

[1,2,4,8,16,32,64,128]
82

*/
