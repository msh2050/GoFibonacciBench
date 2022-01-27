package main

import (
	"math/big"

	godevsample "./goDevSample"
)

func godevsamplefibonacci(n int) *big.Int {
	return godevsample.Getfib(n)
}
