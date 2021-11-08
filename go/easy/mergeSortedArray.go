package main

func merge(nums1 []int, m int, nums2 []int, n int) {

	if m == 0 {
		copy(nums1, nums2)
	} else if n != 0 {

		for i, j := 0, 0; i < m+n; i++ {

			if i == m+j || j == n {
				nums1 = append(nums1[:i], nums2[j:]...)
				break
			} else if nums1[i] > nums2[j] {
				nums1 = append(nums1[:i+1], nums1[i:len(nums1)-1]...)
				nums1[i] = nums2[j]
				j++
			}
		}

	}

}

func main() {

	nums1 := []int{4, 5, 6, 0, 0, 0}
	m := 3
	nums2 := []int{1, 2, 3}
	n := 3

	merge(nums1, m, nums2, n)

}
