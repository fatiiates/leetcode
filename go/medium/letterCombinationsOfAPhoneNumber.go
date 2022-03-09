package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	arr := []string{}
	m := map[byte]string{
		50: "abc",
		51: "def",
		52: "ghi",
		53: "jkl",
		54: "mno",
		55: "pqrs",
		56: "tuv",
		57: "wxyz",
	}
	for i := 0; i < len(digits); i++ {
		if len(arr) == 0 {
			for _, v := range m[digits[i]] {
				arr = append(arr, string(v))
			}
		} else {
			l := len(arr)
			for j := 0; j < l; j++ {
				for _, v := range m[digits[i]] {
					arr = append(arr, arr[j]+string(v))
				}
			}
			arr = arr[l:]
		}
	}

	return arr
}

func main() {
	fmt.Println(letterCombinations("239"))
}