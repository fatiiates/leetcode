package main

func moveZeroes(nums []int) {

	l := len(nums) - 1

	for i := 0; i < l; i++ {
		if nums[i] == 0 {
			nums = append(nums[:i], nums[i+1:]...)
			nums = append(nums, 0)
			l -= 1
		}
	}
}

func main() {

	moveZeroes([]int{})

}
