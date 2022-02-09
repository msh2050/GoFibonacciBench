package main

import "fmt"

func main() {
	type bitlooptests = struct {
		a    int
		want uint
	}

	var bitloop [20]bitlooptests
	for i := 0; i < 20; i++ {

		bitloop[i].a = i + 1
		bitloop[i].want = bitshiftloopbuiltin(i + 1)

		fmt.Printf("{ %v , %v } , ", bitloop[i].a, bitloop[i].want)
	}
	//bitshiftloop
	fmt.Printf("\n\n\n")
	for i := 0; i < 20; i++ {

		bitloop[i].a = i + 1
		bitloop[i].want = bitshiftloop(i + 1)

		fmt.Printf("{ %v , %v } , ", bitloop[i].a, bitloop[i].want)
	}
}
