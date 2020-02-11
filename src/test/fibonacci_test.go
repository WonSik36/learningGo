/*
	recursion vs tail recursion in fibonacci
	https://www.google.com/amp/s/www.geeksforgeeks.org/tail-recursion-fibonacci/amp/
	
	how to test
	terminal:
	 	go test -bench .
*/

package main

import "testing"

func BenchmarkFibonacci50(b *testing.B){
	for i:=0;i<b.N;i++ {
		Fibonacci(50)
	}
}

func Fibonacci(n int) int64 {
	if n==0 {
		return 0
	}else if n==1{
		return 1
	}

	return Fibonacci(n-2) + Fibonacci(n-1)
}

func BenchmarkFibonacciWithMemo50(b *testing.B){
	var memo [100]int64
	for i:=0;i<b.N;i++ {
		FibonacciWithMemo(50, &memo)
	}
}

func FibonacciWithMemo(n int, memo *[100]int64) int64 {
	if n==0 {
		return 0
	}else if n==1{
		return 1
	}

	if memo[n] != 0 {
		return memo[n]
	}else{
		memo[n] = FibonacciWithMemo(n-2, memo) + FibonacciWithMemo(n-1, memo)
		return memo[n]
	}
}

func BenchmarkFibonacciTail50(b *testing.B){
	for i:=0;i<b.N;i++ {
		FibonacciTail(50,0,1)
	}
}

func FibonacciTail(n int, a int64, b int64) int64 {
	if n==0 {
		return a
	}else if n==1{
		return b
	}

	return FibonacciTail(n-1, b, a+b)
}