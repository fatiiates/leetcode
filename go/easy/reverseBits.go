package main

import "fmt"

func reverseBits(num uint32) uint32 {

	var r uint32 = 0
	for i := 0; i < 32; i++ {
		r <<= 1
		if num&1 == 1 {
			r |= 1
		}
		num >>= 1
	}

	return r
}

func main() {

	fmt.Println(reverseBits(434))

}
