package main

import "fmt"

func convertionTable(x byte) int {
	if x == 73 {
		return 1
	}
	if x == 86 {
		return 5
	}
	if x == 88 {
		return 10
	}
	if x == 76 {
		return 50
	}
	if x == 67 {
		return 100
	}
	if x == 68 {
		return 500
	}
	if x == 77 {
		return 1000
	}
	return 0
}

func romanToInt(s string) int {
	result := 0

	for i := 1; i < len(s); i++ {
		if convertionTable(s[i-1]) < convertionTable(s[i]) {
			result += -convertionTable(s[i-1])
		} else {
			result += convertionTable(s[i-1])
		}
	}

	result += convertionTable(s[len(s)-1])

	return result
}

func main() {

	fmt.Println(romanToInt("IIX"))

}
