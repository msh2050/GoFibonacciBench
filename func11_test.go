package main

import (
	"reflect"
	"testing"
)

func Test_fibonacciGoroutine(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fibonacciGoroutine(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("fibonacciGoroutine() = %v, want %v", got, tt.want)
			}
		})
	}
}
