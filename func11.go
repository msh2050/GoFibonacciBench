package main

import (
	"math/big"
	"math/bits"
)

func fibonacciGoroutine(n int) *big.Int {
	if n == 0 {
		return big.NewInt(0)
	}
	chA := make(chan *big.Int)
	chB := make(chan *big.Int)

	num := uint(n)
	mask := bits.RotateLeft(1, bits.Len(num)-1)
	a, b := big.NewInt(0), big.NewInt(1)

	for {
		go CalA(a, b, chA)
		go CalB(a, b, chB)

		a = <-chA
		b = <-chB

		if num&mask > 0 {
			a, b = b, Add(a, b)

		}

		if mask == 1 {
			return a

		}
		mask = bits.RotateLeft(mask, -1)

	}

}
