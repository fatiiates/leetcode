package main

import (
	"fmt"
)

func canConstruct(ransomNote string, magazine string) bool {

	m := make(map[rune]int)

	l := len(ransomNote)
	for _, v := range ransomNote {
		if _, ok := m[v]; ok {
			m[v]++
		} else {
			m[v] = 1
		}
	}
	 
	for _, v := range magazine {
		if _, ok := m[v]; ok {
			
			if m[v] == 1 {
				delete(m, v)
			} else {
				m[v]--
				l--
				if l == 0 {
					return true
				}
			}

			
		}
	}
	 
	if len(m) == 0 {
		return true
	}

	return false
}

func main() {
	fmt.Println(canConstruct("bg", "efjbdfbdgfjhhaiigfhbaejahgfbbgbjagbddfgdiaigdadhcfcj"))
}
