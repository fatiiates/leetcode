package main

import (
	"fmt"
	"strconv"
)

func multiply(num1 string, num2 string) string {

	arr := []string{}
	lenN1 := len(num1) - 1
	lenN2 := len(num2) - 1
	for i := lenN2; i >= 0; i-- {

		n2 := convertAtoi(num2[i])

		c := 0
		localTotal := ""
		for j := lenN1; j >= 0; j-- {
			n1 := convertAtoi(num1[j])
			mul := n1*n2 + c

			if j != 0 {
				rm := mul % 10
				c = mul / 10

				localTotal = strconv.Itoa(rm) + localTotal
			} else {
				localTotal = strconv.Itoa(mul) + localTotal
			}

		}
		if i < lenN2 {
			digits := ""
			for k := i; k < lenN2; k++ {
				digits += "0"
			}
			arr = append(arr, localTotal+digits)
		} else {
			arr = append(arr, localTotal)
		}
	}

	s := sum(&arr)
	i := -1

	for k, v := range s {
		if v != 48 {
			i = k
			break
		}
	}

	if i == -1 {
		return "0"
	}
	return s[i:]
}

func sum(arr *[]string) string {
	m, ls := max(arr)

	total := ""
	c := 0
	for i := 0; i <= m; i++ {
		s := 0
		for k, v := range *arr {
			if ls[k]-i >= 0 && ls[k]-i <= ls[k] {
				n := convertAtoi(v[ls[k]-i])
				s += n
			}
		}
		s += c

		if i != m {
			rm := s % 10
			c = s / 10
			total = strconv.Itoa(rm) + total
		} else {
			total = strconv.Itoa(s) + total
		}

	}
	return total
}

func max(arr *[]string) (int, []int) {
	l := len((*arr)[0]) - 1
	lengths := []int{}
	for _, v := range *arr {
		local := len(v) - 1
		lengths = append(lengths, local)
		if l < local {
			l = local
		}
	}
	return l, lengths
}

func convertAtoi(num uint8) int {
	list := map[int]int{
		48: 0,
		49: 1,
		50: 2,
		51: 3,
		52: 4,
		53: 5,
		54: 6,
		55: 7,
		56: 8,
		57: 9,
	}
	return list[int(num)]
}

func main() {

	fmt.Println(multiply("721", "140"))
}
