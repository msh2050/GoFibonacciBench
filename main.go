package main

import (
	"fmt"
	"math/big"
)

func main() {

	args := []int{100, 1025, 2000, 26587, 4869875, 5987365}

	for _, arg := range args {

		a := fibt(uint(arg))
		b, _ := big.NewInt(0).SetString(fibonaccigmpRC(arg).String(), 10)

		if a.Cmp(b) == 0 {
			fmt.Printf("%v: OK \n", arg)
		} else {
			fmt.Println(a.Sub(a, b))
			fmt.Println(a)
			fmt.Println(b)
			fmt.Printf("%v: FAIL \n", arg)
		}

		fmt.Println()

	}

}
