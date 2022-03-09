package main

import "fmt"

func intToRoman(num int) string {
	roman := ""
	m := map[int]string{
		1:    "I",
		5:    "V",
		10:   "X",
		50:   "L",
		100:  "C",
		500:  "D",
		1000: "M",
	}

	for i := 0; num > 0; num, i = num/10, i+1 {
		romanDigit := ""
		power := pow(10, i)
		digit := num % 10 * power
		if _, ok := m[digit]; ok {
			romanDigit = m[digit]
		} else {
			n := num % 10
			if n == 4 || n == 9 {
				romanDigit = m[power] + m[digit+power]
			} else {
				if n > 5 {
					romanDigit += m[power*5]
					n -= 5
				}
				for i := 0; i < n; i++ {
					romanDigit += m[power]
				}
			}

		}
		roman = romanDigit + roman
	}

	return roman
}

func pow(m, n int) int {
	num := 1
	for i := 0; i < n; i++ {
		num *= m
	}
	return num
}

func main() {
	fmt.Println(intToRoman(57))
	fmt.Println(intToRoman(888))

	fmt.Println(intToRoman(1))
	fmt.Println(intToRoman(2))
	fmt.Println(intToRoman(3))
	fmt.Println(intToRoman(4))
	fmt.Println(intToRoman(5))
	fmt.Println(intToRoman(1000))
	fmt.Println(intToRoman(3999))
	fmt.Println(intToRoman(1001))
	fmt.Println(intToRoman(1453))
	fmt.Println(intToRoman(2023))
}
