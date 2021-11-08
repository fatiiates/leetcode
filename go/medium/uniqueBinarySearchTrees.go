package main

import (
	"fmt"
	"math/big"
)

func numTrees(n int) int {
	var times2 = new(big.Int)
	times2.MulRange(1, int64(2*n))

	var nplus1 = new(big.Int)
	nplus1.MulRange(1, int64(n+1))

	var nth = new(big.Int)
	nth.MulRange(1, int64(n))

	var res = new(big.Int)

	return int(res.Div(times2, res.Mul(nplus1, nth)).Int64())
}

func main() {

	fmt.Println(numTrees(16))
}
