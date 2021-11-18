package main

import "fmt"

func letterCasePermutation(s string) []string {
	return permutation(s, "", &[]string{})
}

func permutation(nums string, local string, global *[]string) []string {
	l := len(nums)

	if l == 0 {
		*global = append(*global, local)
		return *global
	}

	if nums[0] > 96 {
		permutation(nums[1:], local+string(nums[0]-32), global)
	} else if nums[0] > 64 {
		permutation(nums[1:], local+string(nums[0]+32), global)
	}
	local += string(nums[0])
	permutation(nums[1:], local, global)

	return *global
}

func main() {

	fmt.Println(letterCasePermutation("a1b2c3d4e5f6"))
	fmt.Println(letterCasePermutation("1a2b3c4d5e6f"))
}
