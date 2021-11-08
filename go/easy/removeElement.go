package main

import "fmt"

func removeElement(nums []int, val int) int {
	i := 0
	for i < len(nums) {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			i -= 1
		}
		i += 1
	}

	return i
}

func main() {

	fmt.Println(removeElement([]int{0, 1, 2, 2, 3, 0, 4, 2}, 2))

}
