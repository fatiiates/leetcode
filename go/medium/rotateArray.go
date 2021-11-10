package main

func rotate(nums []int, k int) {

	l := len(nums)
	m := k % l
	t := append(nums[l-m:], nums[:l-m]...)
	copy(nums, t)
}

func main() {

	rotate([]int{1, 2, 3, 4, 5, 6, 7}, 3)
}
