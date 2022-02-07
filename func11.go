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
	chB1 := make(chan *big.Int)
	chB2 := make(chan *big.Int)

	num := uint(n)
	mask := bits.RotateLeft(1, bits.Len(num-1))
	a, b := big.NewInt(0), big.NewInt(1)

	for mask > 0 {
		go calA(a, b, chA)
		go calB1(a, chB1)
		go calB2(b, chB2)

		a = <-chA
		b = Add(<-chB1, <-chB2)

		if num&mask > 0 {
			a, b = b, Add(a, b)

		}

		mask /= 2

	}

	return a
}

func calA(a, b *big.Int, ch chan *big.Int) {

	ch <- Mul(a, Sub(Mul(b, big.NewInt(2)), a))
}

/*func calD(a, b *big.Int, ch chan *big.Int) {
	ch <- Add(Mul(a, a), Mul(b, b))
}*/
func calB1(a *big.Int, ch chan *big.Int) {

	ch <- Mul(a, a)
}
func calB2(b *big.Int, ch chan *big.Int) {

	ch <- Mul(b, b)
}
