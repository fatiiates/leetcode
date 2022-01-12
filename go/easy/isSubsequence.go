package main

import "fmt"

func isSubsequence(s string, t string) bool {
	if len(s) == 0 || len(t) == 0 {
		return false
	}
	for _, vs := range s {
		j := 0
		for _, vt := range t {
			fmt.Println(vt, t, vs, s)
			if vs == vt {
				s = s[1:]
				fmt.Println("match")
				break
			}
			j++
		}
		if len(s) != 0 && len(t) == j {
			return false
		}
		fmt.Println(t, s, j, len(t))
		t = t[j+1:]
	}
	if len(s) == 0 {
		return true
	}
	return false
}

func main() {
	fmt.Println(isSubsequence("abb", "abcb"))
}
