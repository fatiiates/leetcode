package main

import "fmt"

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, v := range nums1 {
		if _, ok := m[v]; ok {
			m[v] += 1
		} else {
			m[v] = 1
		}
	}

	arr := []int{}

	for _, v := range nums2 {
		if val, ok := m[v]; ok && val > 0 {
			m[v] -= 1
			arr = append(arr, v)
		}
	}

	return arr
}

func main() {

	fmt.Println(intersect([]int{4, 9, 5}, []int{9, 4, 9, 8, 4}))

}
