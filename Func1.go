package main

import (
	"fmt"
	"math/big"
)

func main() {

	for i := 0; i <= 10; i++ {

		fmt.Printf("The %v:th fibonacci number is %v\n", i, fibonacci(i))
	}

	n := 1001

	fmt.Printf("The %v:th fibonacci number is %v\n", n, fibonacci(n))

	n = 500 * 1000

	fmt.Printf("The %v:th fibonacci number has %v digits\n", n, len(fibonacci(n).String()))

}

func fibonacci(n int) *big.Int {
	fn := make(map[int]*big.Int)
	fn[0] = big.NewInt(0)
	fn[1] = big.NewInt(1)

	for i := 2; i <= n; i++ {
		var f = big.NewInt(0)
		f = f.Add(fn[i-1], fn[i-2])
		fn[i] = f
	}
	return fn[n]
}
