package main

import (
	nouse "github.com/msh2050/GoFibonacciBench/nouseFibgmp"
	big "github.com/msh2050/gmp"
)

func nousefibonaccigmp(n int) *big.Int {

	return nouse.FastDoubling(n)
}
