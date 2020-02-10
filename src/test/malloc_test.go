/*
	memory allocation

	how to test
	terminal:
	 	go test -bench . -benchmem
*/

package main

import (
	"testing"
)

var global *int
func BenchmarkMalloc(b *testing.B) {
	// if not using benchmem option than using this function
	// b.ReportAllocs()
	for i:=0;i<b.N;i++ {
		global = new(int)
	}
}