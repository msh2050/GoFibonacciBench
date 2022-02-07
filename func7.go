package main

import (
	gmp "github.com/msh2050/gmp"
	"math/big"
)

func fibonaccigmp(n int) *big.Int {
	f := gmp.NewInt(0)
	ff := f.Fibui(n)

	gb := new(big.Int)
	gb = gb.SetBytes(ff.Bytes())

	return gb
}
func fibonaccigmpRC(n int) *gmp.Int {
	f := gmp.NewInt(0)
	return f.Fibui(n)
}
