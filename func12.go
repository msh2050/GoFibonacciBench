package main

import (
	"math/big"
	"math/bits"
)

//https://chunminchang.github.io/blog/post/calculating-fibonacci-numbers-by-fast-doubling
func fibonaccichanggoroutin(n uint) *big.Int {
	if n == 0 {
		return big.NewInt(0)
	}

	// num := uint(n)
	chA := make(chan *big.Int)
	chB := make(chan *big.Int)
	a, b := big.NewInt(0), big.NewInt(1)

	// There is only one `1` in the bits of `mask`. The `1`'s position is same as
	// the highest bit of n(mask = 2^(h-1) at first), and it will be shifted right
	// iteratively to do `AND` operation with `n` to check `n_j` is odd or even,
	// where n_j is defined below.
	mask := bits.RotateLeft(1, bits.Len(uint(n)-1))
	for {

		go CalA(a, b, chA)
		go CalB(a, b, chB)

		c := <-chA // F(2k) = F(k) * [ 2 * F(k+1) – F(k) ]
		d := <-chB // F(2k+1) = F(k)^2 + F(k+1)^2

		if n&mask > 0 { // n_j is odd: k = (n_j-1)/2 => n_j = 2k + 1
			a = d         //   F(n_j) = F(2k + 1)
			b = Add(c, d) //   F(n_j + 1) = F(2k + 2) 1= F(2k) + F(2k + 1)

		} else { // n_j is even: k = n_j/2 => n_j = 2k
			a = c //   F(n_j) = F(2k)
			b = d //   F(n_j + 1) = F(2k + 1)
		}
		if mask == 1 {
			return a

		}
		mask = bits.RotateLeft(mask, -1)
	}

}
