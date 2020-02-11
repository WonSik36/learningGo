/*
	recursion vs tail recursion in factorial
	https://blog.usejournal.com/tail-recursion-in-go-fb5cf69a0f26
	https://aidanbae.github.io/code/golang/benchmark/
	http://pyrasis.com/book/GoForTheReallyImpatient/Unit61/01

	how to test
	terminal:
	 	go test -bench .
*/

package main

import (
	"testing"
)

func BenchmarkFactorial(b *testing.B) {
	for i:=0;i<b.N;i++ {
		Factorial(20)
	}
}

func BenchmarkFactorialTail(b *testing.B) {
	for i:=0;i<b.N;i++ {
		FactorialTail(20,1)
	}
}

// tail recurstion
func FactorialTail(n int64, res int64) int64 {
	if n==1{
		return res
	}

	return FactorialTail(n-1, res*n)
}

// recursion
func Factorial(n int64) int64 {
	if n==1 {
		return 1
	}

	return Factorial(n-1)*n
}