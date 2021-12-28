package main

import "fmt"

func isAnagram(s string, t string) bool {
	sLen := len(s)
	tLen := len(t)
	if sLen != tLen {
		return false
	}
	var sMap [26]int
	var tMap [26]int
	//sMap := make([]int, 26)
	for i := 0; i < sLen; i++ {
		sMap[int(s[i]%97)]++
		tMap[int(t[i]%97)]++
	}

	for i := 0; i < sLen; i++ {
		sMap[int(s[i]%97)]++
		tMap[int(t[i]%97)]++
	}

	return sMap == tMap
}

func main() {

	fmt.Println(isAnagram("anagram", "aganarm"))

}
