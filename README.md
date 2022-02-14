# GoFibonacciBench

Benchmarking many Fibonacci functions and algorithms with running unit test ...

## Introduction:
Fibonacci's numbers are special kinds of numbers that form a special sequence.
This sequence is one of the famous formulas in mathematics.
You can find Fibonacci numbers in plant and animal structures.
These numbers are also called nature's universal rule, and nature's secret code.

Fibonacci's numbers are a sequence of whole numbers arranged as:-

```
0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, ...
```

## what we benchmark
there are so many code in the net to get Fibonacci number, but I think I can do better

## method
1.  collect the methods used to get Fibonacci in go  :white_check_mark:
2. run unit test to check and correct (some function start form 1 not 0) :white_check_mark:
3. check function limitation :white_check_mark:
4. benchmark :white_check_mark:
5. try to make/modify/enhance for best results :white_check_mark:


## Results:- ##
![alt text](https://github.com/msh2050/GoFibonacciBench/blob/main/plot/operations.png?raw=true)

As in the table above we made our best results with the last function (Func14).
Some function excluded from the table because it will be very slow in big number because they use iteration method.
The fast doubling is the fastest reliable method ... 
The Binet formula will be faster, but it will not be accurate since we have to use floating point.
There are many tricks to increase the speed.
The main one is using Goroutine to run two equation concurrently.
the other is to use table for the first 1025 fib numbers.
the table will take about 200KB. will be used to get the fib number directly for the first 1025.
the other use of the table is to remove the first 10 iteration of the fast doubling(2^10 =1024).

++++++++++++++++++++++++++++++++++++++++++++++++++++++++

## sources:-
- <https://www.cuemath.com/algebra/fibonacci-numbers/>

- <https://github.com/T-PWK/go-fibonacci>
- <https://stackoverflow.com/questions/33664041/how-to-convert-int-to-bigint-in-golang>
- <https://www.nayuki.io/page/fast-fibonacci-algorithms>
- https://gist.github.com/parasyte/4077694
- [Fibonacci benchmark correctly in Golang and Ruby](https://gist.github.com/nouse/192029c913edacebb39e4859449d1180) 
- https://go.dev/test/fibo.go