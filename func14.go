package main

import (
	"math/big"
	"math/bits"
)

//https://chunminchang.github.io/blog/post/calculating-fibonacci-numbers-by-fast-doubling
func fibt(n uint) *big.Int {
	if n < 1025 {
		return big.NewInt(0).SetBytes(Fibtable[n])
	}
	chA := make(chan *big.Int)
	chB := make(chan *big.Int)
	a, b := big.NewInt(0), big.NewInt(1)

	// There is only one `1` in the bits of `mask`. The `1`'s position is same as
	// the highest bit of n(mask = 2^(h-1) at first), and it will be shifted right
	// iteratively to do `AND` operation with `n` to check `n_j` is odd or even,
	// where n_j is defined below.

	mask := bits.RotateLeft(1, bits.Len(n)-1)

	//we can ship the first 10 iterations of Doubling fibonacci numbers and get them from the table
	a = big.NewInt(0).SetBytes(Fibtable[n>>(bits.Len(mask)-10)])
	b = big.NewInt(0).SetBytes(Fibtable[n>>(bits.Len(mask)-10)+1])

	mask = bits.RotateLeft(mask, -10)
	for {

		if mask == 1 {

			// we are on last iteration wee need to just calculate a
			if n&mask > 0 { // n_j is odd: k = (n_j-1)/2 => n_j = 2k + 1

				a = Add(Mul(a, a), Mul(b, b))

			} else { // n_j is even: k = n_j/2 => n_j = 2k

				a = Mul(a, Sub(Mul(b, big.NewInt(2)), a))

			}

			return a

		}

		go CalA(a, b, chA)
		go CalB(a, b, chB)

		if n&mask > 0 { // n_j is odd: k = (n_j-1)/2 => n_j = 2k + 1

			a = <-chB
			b = Add(a, <-chA)
		} else { // n_j is even: k = n_j/2 => n_j = 2k

			a = <-chA
			b = <-chB
		}

		mask = bits.RotateLeft(mask, -1)

	}

}
