package main

import (
	"fmt"
	"testing"
)

var bittests = []struct {
	a    int
	want uint
}{
	{1, 1},
	{2, 2},
	{3, 2},
	{4, 4},
	{5, 4},
	{6, 4},
	{7, 4},
	{8, 8},
	{9, 8},
	{16, 16},
	{1560, 1024},
	{1000786, 524288},
}

func Test_bitshiftmath(t *testing.T) {

	for _, tt := range bittests {

		testname := fmt.Sprintf("%d", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := bitshiftmath(tt.a)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

//bitshiftbuildin
func Test_bitshiftbuildin(t *testing.T) {

	for _, tt := range bittests {

		testname := fmt.Sprintf("%d", tt.a)
		t.Run(testname, func(t *testing.T) {
			ans := bitshiftbuildin(tt.a)
			if ans != int(tt.want) {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
func Test_bitshiftloop1(t *testing.T) {
	for i := 1; i > 20; i++ {
		if (bitshiftloop(i)) != (bitshiftloopbuiltin(i)) {
			t.Errorf("bitshiftloop : %d, bitshiftloopbuild %d", bitshiftloop(i), bitshiftloopbuiltin(i))
		}
	}
}

func Benchmark_bitshiftmath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bitshiftmath(1000000)
	}
}

//bitshiftbuildin
func Benchmark_bitshiftbuildin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bitshiftbuildin(1000000)
	}
}

//bitshiftloop
func Benchmark_bitshiftloop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bitshiftloop(1000000)
	}
}

//bitshiftloopbuiltin
func Benchmark_bitshiftloopbuiltin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bitshiftloopbuiltin(1000000)
	}
}
