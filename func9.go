// Borthralla fibonacci implementation: https://github.com/golang/go/issues/30943
package main

import (
	"fmt"
	"math/big"
)

func fibBorthralla(n int) *big.Int {
	if n == 0 {
		return big.NewInt(0)
	}
	left := big.NewInt(1)
	right := big.NewInt(1)

	helper1 := big.NewInt(0)
	helper2 := big.NewInt(0)

	bin := fmt.Sprintf("%b", n)
	for i := 1; i < len(bin); i++ {

		helper1.Mul(big.NewInt(2), right)
		helper1.Sub(helper1, left)
		helper1.Mul(left, helper1)

		helper2.Mul(right, right)
		left.Mul(left, left)
		helper2.Add(helper2, left)

		if bin[i] == '0' {
			left.Set(helper1)
			right.Set(helper2)
		} else {
			left.Set(helper2)
			right.Add(helper1, helper2)
		}
	}

	return left
}
