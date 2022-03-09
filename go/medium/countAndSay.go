package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	say := "1"
	for i := 1; i < n; i++ {
		newSay := ""
		count := 0
		for j := 0; j < len(say); j++ {
			if j < len(say)-1 && say[j] == say[j+1] {
				count++
			} else {
				count++
				newSay += strconv.Itoa(count) + string(say[j])
				count = 0
			}
		}
		fmt.Println(newSay)
		say = newSay
	}
	return say
}

func main() {
	fmt.Println(countAndSay(4))
}
