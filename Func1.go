package main

import (
	"fmt"
	"math/big"
)

func main() {

	fmt.Printf("fib: %v", fibonacciGoroutine(100))

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
