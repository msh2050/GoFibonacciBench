package main

import (
	"math/big"
	"math/bits"
)

func fibonaccibinary(n int) *big.Int {
	if n == 0 {
		return big.NewInt(0)
	}
	num := uint(n)
	mask := bits.RotateLeft(1, bits.Len(num-1))
	a, b := big.NewInt(0), big.NewInt(1)

	for mask > 0 {

		c := Mul(a, Sub(Mul(b, big.NewInt(2)), a))

		d := Add(Mul(a, a), Mul(b, b))
		a = c
		b = d

		if num&mask > 0 {
			a, b = b, Add(a, b)

		}

		mask /= 2

	}

	return a
}
