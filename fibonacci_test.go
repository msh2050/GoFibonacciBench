package main

import (
	"fmt"
	"math/big"
	"reflect"
	"testing"
)

type args struct {
	n int
}

var table = []struct {
	input int
}{
	{input: 100},
	{input: 1000},
	{input: 50000},
	{input: 100000},
	{input: 1000000},
}
var fibof100, _ = new(big.Int).SetString("354224848179261915075", 10)
var Fibof1000, _ = new(big.Int).SetString("43466557686937456435688527675040625802564660517371780402481729089536555417949051890403879840079255169295922593080322634775209689623239873322471161642996440906533187938298969649928516003704476137795166849228875", 10)

var tests = []struct {
	name string
	args args
	want *big.Int
}{
	{"number1(0)", args{0}, big.NewInt(0)}, {"number2(1)", args{1}, big.NewInt(1)},
	{"number3(2)", args{2}, big.NewInt(1)}, {"number4(3)", args{3}, big.NewInt(2)},
	{"number5(4)", args{4}, big.NewInt(3)}, {"number6(5)", args{5}, big.NewInt(5)},
	{"number7(6)", args{6}, big.NewInt(8)}, {"number8(7)", args{7}, big.NewInt(13)},
	{"number9(8)", args{8}, big.NewInt(21)}, {"number10(9)", args{9}, big.NewInt(34)},
	{"number11(100)", args{100}, fibof100}, {"number12(1000)", args{1000}, Fibof1000},
}

func Test_fibonacci(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacci(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fibonaccifast(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonaccifast(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fibonaccifast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_tpwkfibonacci(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tpwkfibonacci(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tpwkfibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

//fibonaccibinary
func Test_fibonaccibinary(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonaccibinary(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("tpwkfibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

// godevsamplefibonacci
func Test_godevsamplefibonacci(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := godevsamplefibonacci(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("godevsamplefibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

//nousefibonacci
func Test_nousefibonacci(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := nousefibonacci(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("godevsamplefibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

// fibonaccigmp
func Test_fibonaccigmp(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonaccigmp(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("godevsamplefibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

//nousefibonaccigmp
func Test_nousefibonaccigmp(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gmpgot := nousefibonaccigmp(tt.args.n)
			aaa := gmpgot.Bytes()
			bbb := new(big.Int)
			bbb = bbb.SetBytes(aaa)
			if got := bbb; !reflect.DeepEqual(got, tt.want) {
				t.Errorf("godevsamplefibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

//fibBorthralla
func Test_fibBorthralla(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibBorthralla(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("godevsamplefibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

//fibonacciGoroutine func11
func Test_fibonacciGoroutine(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacciGoroutine(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("godevsamplefibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

//fibonaccichanggoroutin
func Test_fibonaccichanggoroutin(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonaccichanggoroutin(uint(tt.args.n)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("godevsamplefibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

//fibonaccichang
func Test_fibonaccichang(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonaccichang(uint(tt.args.n)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("godevsamplefibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

//fibt
func Test_fibt(t *testing.T) {

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibt(uint(tt.args.n)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("godevsamplefibonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}

// bench mark for fibonacci(Func1)
func Benchmark_fibonacci(b *testing.B) {

	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fibonacci(v.input)
			}
		})
	}
}

// bench mark for fibonaccifast(Func2)
func Benchmark_fibonaccifast(b *testing.B) {

	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fibonaccifast(v.input)
			}
		})
	}
}

// bench mark for tpwkfibonacci (Func3)
func Benchmark_tpwkfibonacci(b *testing.B) {

	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				tpwkfibonacci(v.input)
			}
		})
	}
}

//
// bench mark for godevsamplefibonacci (Func4)
func Benchmark_godevsamplefibonacci(b *testing.B) {

	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				godevsamplefibonacci(v.input)
			}
		})
	}
}

//
// bench mark for nousefibonacci (Func5)
func Benchmark_nousefibonacci(b *testing.B) {

	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				nousefibonacci(v.input)
			}
		})
	}
}

// bench mark for fibonaccibinary (Func6)
func Benchmark_fibonaccibinary(b *testing.B) {

	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fibonaccibinary(v.input)
			}
		})
	}
}

//fibonaccigmp func 7
func Benchmark_fibonaccigmp(b *testing.B) {

	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fibonaccigmp(v.input)
			}
		})
	}
}

//nousefibonaccigmp func8
func Benchmark_nousefibonaccigmp(b *testing.B) {

	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				nousefibonaccigmp(v.input)
			}
		})
	}
}

//fibBorthralla func9
func Benchmark_fibBorthralla(b *testing.B) {

	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fibBorthralla(v.input)
			}
		})
	}
}

//fibBorthrallagmp func10
func Benchmark_fibBorthrallagmp(b *testing.B) {

	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fibBorthrallagmp(v.input)
			}
		})
	}
}

//fibonaccigmpRC func7 gmp
func Benchmark_fibonaccigmpRC(b *testing.B) {

	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fibonaccigmpRC(v.input)
			}
		})
	}
}

//fibonacciGoroutine func11
func Benchmark_fibonacciGoroutine(b *testing.B) {

	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fibonacciGoroutine(v.input)
			}
		})
	}
}

//fibonaccichanggoroutin
func Benchmark_fibonaccichanggoroutin(b *testing.B) {

	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fibonaccichanggoroutin(uint(v.input))
			}
		})
	}
}

//fibonaccichang
func Benchmark_fibonaccichang(b *testing.B) {

	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fibonaccichang(uint(v.input))
			}
		})
	}
}

//fibt
func Benchmark_fibt(b *testing.B) {

	for _, v := range table {
		b.Run(fmt.Sprintf("input_size_%d", v.input), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				fibt(uint(v.input))
			}
		})
	}
}
