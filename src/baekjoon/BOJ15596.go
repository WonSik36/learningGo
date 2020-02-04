/*
	baekjoon online judge
	problem number 15596
	https://www.acmicpc.net/problem/15596
*/

package main

import (
	"fmt"
)

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sum(arr))
}

func sum(a []int) int {
	sum := 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}

	return sum
}
