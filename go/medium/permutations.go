package main

import "fmt"

func permute(nums []int) [][]int {
	return permutation(nums, []int{}, &[][]int{})
}

func permutation(nums []int, local []int, global *[][]int) [][]int {
	l := len(nums)

	if l == 0 {
		arr := make([]int, len(local))
		copy(arr, local)
		*global = append(*global, arr)
		return *global
	}

	for i := 0; i < l; i++ {
		local = append(local, nums[i])
		lnums := make([]int, l)
		copy(lnums, nums)
		lnums = append(lnums[:i], lnums[i+1:]...)
		permutation(lnums, local, global)
		local = local[0 : len(local)-1]
	}

	return *global
}

func main() {

	fmt.Println(permute([]int{-10, -9, -8, 8, 9, 10}))
}
