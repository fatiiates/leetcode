package main

import "fmt"

func singleNonDuplicate(nums []int) int {
	l := len(nums)
	if l == 1 {
		return nums[0]
	}
	return travel(&nums, 0, l, l-1)
}

func travel(nums *[]int, s int, e int, mx int) int {
	if s > e {
		return -1
	}
	if e-s == 0 {

		if s == mx {
			if (*nums)[s-1] != (*nums)[s] {
				return (*nums)[s]
			}
		} else if s == 0 {
			if (*nums)[s] != (*nums)[s+1] {
				return (*nums)[s]
			}
		} else if (*nums)[s-1] != (*nums)[s] && (*nums)[s] != (*nums)[s+1] {
			return (*nums)[s]
		}
	} else {

		p := (e - s) / 2
		t1 := travel(nums, s, s+p, mx)
		if t1 != -1 {
			return t1
		}
		t2 := travel(nums, s+p+1, e, mx)
		if t2 != -1 {
			return t2
		}
	}
	return -1
}

func main() {

	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 2, 3}))
	fmt.Println(singleNonDuplicate([]int{1, 1, 2, 3, 3}))
	fmt.Println(singleNonDuplicate([]int{1, 2, 2, 3, 3}))
}
