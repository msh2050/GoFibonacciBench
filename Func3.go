package main

import (
	"math/big"

	tpwk "github.com/t-pwk/go-fibonacci"
)

/*
implementing from
High-performance Fibonacci numbers implementation in Go
https://github.com/T-PWK/go-fibonacci
*/

func tpwkfibonacci(n int) *big.Int {
	return tpwk.FibonacciBig(uint(n))
}
