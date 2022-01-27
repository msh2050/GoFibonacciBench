package main

import (
	"math/big"

	godevsample "github.com/msh2050/GoFibonacciBench/goDevSample"
)

func godevsamplefibonacci(n int) *big.Int {
	return godevsample.Getfib(n)
}
