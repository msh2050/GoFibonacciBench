// Package fib source https://gist.github.com/https://gist.github.com/nouse/192029c913edacebb39e4859449d1180/192029c913edacebb39e4859449d1180
package fib

import (
	"testing"

	big "github.com/ncw/gmp"
)

func iterateAdding(n int) *big.Int {
	a := big.NewInt(0)
	b := big.NewInt(1)

	for i := 0; i < n; i++ {
		a.Add(a, b)
		a, b = b, a
	}
	return a
}

func FastDoubling(n int) string {
	a, _, _ := _fastDoubling(n)
	return a.String()
}

// (Private) Returns the tuple (F(n), F(n+1)).
func _fastDoubling(n int) (*big.Int, *big.Int, *big.Int) {
	if n == 0 {
		return big.NewInt(0), big.NewInt(1), new(big.Int)
	}

	a, b, c := _fastDoubling(n / 2)

	// (2*b - a)*a
	c.Mul(a, b)
	c.Lsh(c, 1)

	a.Mul(a, a)
	c.Sub(c, a)
	// a*a + b*b
	b.Mul(b, b)
	b.Add(b, a)
	if n%2 == 0 {
		return c, b, a
	}

	a.Add(c, b)
	return b, a, c
}

func BenchmarkIterateAdding(b *testing.B) {
	for i := 0; i < b.N; i++ {
		iterateAdding(100000)
	}
}

func BenchmarkFastDoubling(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FastDoubling(100000)
	}
}
