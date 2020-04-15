/*
	baekjoon online judge
	problem number 15735
	https://www.acmicpc.net/problem/15735

	Greedy Algorithm
	1. 2N-th floor triangle has
	sum(1~2N) + 2*sum(1~(2N-1)) + sum(1~(2N-2)) + 2*sum(1~(2N-3)) + ... + sum(1~2) + 2*sum(1)
	2. (2N-1)-th floor triangle has
	sum(1~(2N-1)) + 2*sum(1~(2N-2)) + sum(1~(2N-3)) + sum(1~(2N-4)) + ... + 2*sum(1~2) + sum(1)
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	N := getInt(r)
	sum := make([]int64, N+1, N+1)

	sum[0] = 0
	for i := 1; i <= N; i++ {
		sum[i] += sum[i-1] + int64(i)
	}

	var constant int64 = 1
	var result int64 = 0
	for i := N; i > 0; i-- {
		result += constant * sum[i]

		if constant%2 == 0 {
			constant = 1
		} else {
			constant = 2
		}
	}

	w.WriteString(strconv.FormatInt(result, 10))
	w.Flush()
}

func getInt(r *bufio.Reader) int {
	res, _ := strconv.Atoi(getString(r))

	return res
}

func getString(r *bufio.Reader) string {
	str, _ := r.ReadString('\n')
	str = strings.TrimSpace(str)

	return str
}

func getStrTkn(r *bufio.Reader) []string {
	str, _ := r.ReadString('\n')
	str = strings.TrimSpace(str)
	strTkn := strings.Split(str, " ")

	return strTkn
}

func getIntArr(r *bufio.Reader) []int {
	strTkn := getStrTkn(r)
	arr := make([]int, len(strTkn), len(strTkn))

	for i := 0; i < len(arr); i++ {
		arr[i], _ = strconv.Atoi(strTkn[i])
	}

	return arr
}

func printArr(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Print("\n")
}
