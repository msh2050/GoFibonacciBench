package main

import (
	"math/bits"
)

func bitshiftmath(n int) uint {

	mask := bits.RotateLeft(1, bits.Len(uint(n))-1)

	return mask
}

//unsigned int h = 0;
//for (unsigned int i = n ; i ; ++h, i >>= 1);
//unsigned int mask = 1 << (h - 1)
func bitshiftbuildin(n int) int {

	h := 0
	for i := n; i > 0; i >>= 1 {
		h++
	}
	mask := 1 << (h - 1)

	return mask
}

func bitshiftloop(n int) uint {

	a := uint(0)
	mask := bits.RotateLeft(1, n)
	for {
		a++
		if mask == 1 {
			return a
		}
		mask = bits.RotateLeft(mask, -1)
	}

}

func bitshiftloopbuiltin(n int) uint {

	a := uint(0)
	for mask := bits.RotateLeft(1, n); mask > 0; mask >>= 1 {
		a++
	}
	return a
}
