package main

import "fmt"

func maxProduct(nums []int) int {

	mx := nums[0]
	mn := nums[0]
	res := nums[0]

	l := len(nums)
	for i := 1; i < l; i++ {

		tmp := mx

		mx = max(max(mx*nums[i], mn*nums[i]), nums[i])
		mn = min(min(tmp*nums[i], mn*nums[i]), nums[i])

		if mx > res {
			res = mx
		}
	}

	return res

}

func max(n1 int, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}

func min(n1 int, n2 int) int {
	if n1 < n2 {
		return n1
	}
	return n2
}

func main() {
	fmt.Println([]int{-2, 3, -4})
}
