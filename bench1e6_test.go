package main

import (
	"math/big"
	"reflect"
	"testing"
)

var test1x6 = []struct {
	name string
	arg  int
	want *big.Int
}{
	{"number12(1000)", 1000, Fibof1000},
}

func Test_func2(t *testing.T) {

	t.Run(test1x6[0].name, func(t *testing.T) {
		if got := fibonacciGoroutine(test1x6[0].arg); !reflect.DeepEqual(got, test1x6[0].want) {
			t.Errorf("fibonacciGoroutine() = %v, want %v", got, test1x6[0].want)
		}
	})
}
func BenchmarkFunc2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonaccifast(1000000)
	}
}

func BenchmarkFunc5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nousefibonacci(1000000)
	}
}

func BenchmarkFunc6(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonaccibinary(1000000)
	}
}

func BenchmarkFunc7GMP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonaccigmpRC(1000000)
	}
}

func BenchmarkFunc8GMP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		nousefibonaccigmp(1000000)
	}
}

func BenchmarkFunc9(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibBorthralla(1000000)
	}
}

func BenchmarkFunc10GMP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibBorthrallagmp(1000000)
	}
}

func BenchmarkFunc11(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciGoroutine(1000000)
	}
}
func BenchmarkFunc12(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacciGoroutine(1000000)
	}
}
func BenchmarkFunc13(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonaccichang(1000000)
	}
}
func BenchmarkFunc14(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibt(1000000)
	}
}

//go test -bench=^BenchmarkFunc . -run=^$ -benchtime=100x -benchmem
