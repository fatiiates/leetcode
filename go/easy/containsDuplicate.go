package main

import "fmt"

func containsDuplicate(nums []int) bool {

	m := make(map[int]bool)

	for _, v := range nums {
		fmt.Println(v)
		if _, ok := m[v]; ok {
			return true
		} else {
			m[v] = true
		}
	}

	return false
}

func main() {

	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))

}
