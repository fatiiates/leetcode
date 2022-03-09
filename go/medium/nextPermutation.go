package main

import "fmt"

func nextPermutation(nums []int) {
	j := -1
	l := len(nums)
	for i := l - 1; i > 0; i-- {
		if nums[i] > nums[i-1] {
			j = i - 1
			break
		}
	}
	
	num := j
	i := j+1
	if j == -1 { 
		i = l
	}
	for ; i < l; i++ {
		if nums[i] > nums[j] {
			num = i
		}
	}
	if num != j {
		nums[num], nums[j] = nums[j], nums[num]
	}
	reverse(&nums, j)
}

func reverse(nums *[]int, j int) {
	l := len(*nums)-1
	for i := 0; i < (l-j)/2; i++ {
		(*nums)[j+1+i], (*nums)[l-i] = (*nums)[l-i], (*nums)[j+1+i]
	}
}

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Println(nums) // 1,3,2

	nums = []int{3,2,1}
	nextPermutation(nums) 
	fmt.Println(nums) // 1,2,3


	nums = []int{1,1,5} 
	nextPermutation(nums) 
	fmt.Println(nums) // 1,5,1

	nums = []int{0} 
	nextPermutation(nums) 
	fmt.Println(nums) // 0

	nums = []int{0,0}
	nextPermutation(nums)
	fmt.Println(nums) // 1,3,2

	nums = []int{1, 2, 3,4}
	nextPermutation(nums)
	fmt.Println(nums) // 1,3,2
}
