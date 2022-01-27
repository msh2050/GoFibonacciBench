package main

import (
	nouse "github.com/msh2050/GoFibonacciBench/nouseFib"
	"math/big"
)

func nousefibonacci(n int) *big.Int {

	return nouse.FastDoubling(n)
}
