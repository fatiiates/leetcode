package main

import (
	"fmt"
)

func dailyTemperatures(temperatures []int) []int {

	l := len(temperatures)
	for i := 0; i < l; i++ {
		ctrl := true
		for j := 1; j < l-i; j++ {
			if temperatures[i] < temperatures[i+j] {
				temperatures[i] = j
				ctrl = false
				break
			}
		}
		if ctrl {
			temperatures[i] = 0

		}

	}

	return temperatures
}

func main() {

	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))

}
