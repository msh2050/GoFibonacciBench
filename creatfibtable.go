package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}
func creatfibtable(n int) {
	f, err := os.Create("fibtable/fibtable.txt")
	check(err)
	defer func(f *os.File) {
		err := f.Close()
		check(err)
	}(f)

	w := bufio.NewWriter(f)
	for i := 0; i <= n; i++ {

		a := fibonacciGoroutine(i).Bytes()
		semiformat := fmt.Sprintf("%v", a)
		semiformat = strings.ReplaceAll(semiformat, "[", "{")
		semiformat = strings.ReplaceAll(semiformat, "]", "}")
		tokens := strings.Split(semiformat, " ")
		_, err = fmt.Fprint(w, strings.Join(tokens, ", "))
		check(err)
		_, err = fmt.Fprintf(w, " , ")
		check(err)

	}

	err = w.Flush()
	check(err)
	fmt.Println("Fibtable created")
}
