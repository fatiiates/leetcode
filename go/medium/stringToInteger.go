package main

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi(s string) int {
	num := 0
	s = strings.Trim(s, " ")
	sign := 1
	limit := math.MaxInt32
	if len(s) < 1 {
		return 0
	}
	if s[0] == '-' {
		sign = -1
		s = s[1:]
		limit = math.MinInt32
	} else if s[0] == '+' {
		s = s[1:]
	}
	s = strings.TrimLeft(s, "0")

	arr := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] >= 48 && s[i] <= 57 {
			arr = append(arr, s[i])
			if i > 9 {
				return limit
			}
		} else {
			break
		}
	}

	for i, v := range arr {
		digit := int(v-48) * pow(10, len(arr)-i-1)
		if num+digit > limit*sign {
			return limit
		}
		num += digit
	}

	return num*sign
}

func pow(m, n int) int {
	num := 1
	for i := 0; i < n; i++ {
		num *= m
	}
	return num
}

func main() {
	fmt.Println(myAtoi("-0000002147483649"))
	fmt.Println(myAtoi("-0000002147483648"))
	fmt.Println(myAtoi("-0000002147483647"))

	fmt.Println(myAtoi("0000002147483646"))
	fmt.Println(myAtoi("0000002147483647"))
	fmt.Println(myAtoi("0000002147483648"))

}
