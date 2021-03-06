package main

import (
	"fmt"
)

type Node struct {
	arr  []int
	nums []int
}

func permute(nums []int) [][]int {
	res := [][]int{}
	if len(nums) == 1 {
		return [][]int{{1}}
	}
	bt := []Node{}
	for i := 0; i < len(nums); i++ {
		lnums := make([]int, len(nums))
		copy(lnums, nums)
		lnums = append(lnums[:i], lnums[i+1:]...)

		bt = append(bt, Node{
			arr:  []int{nums[i]},
			nums: lnums,
		})
	}

	for i := 0; i < len(bt); i++ {
		node := bt[i]
		if len(node.nums) == 0 {
			res = append(res, append([]int{}, node.arr...))
			continue
		}
		for j := 0; j < len(node.nums); j++ {
			lnums := make([]int, len(node.nums))
			copy(lnums, node.nums)
			lnums = append(lnums[:j], lnums[j+1:]...)
			bt = append(bt, Node{
				arr:  append(append([]int{}, node.arr...), node.nums[j]),
				nums: lnums,
			})
		}
	}

	return res
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	// fmt.Println(permute([]int{-10, -9, -8, 8, 9, 10}))
}
