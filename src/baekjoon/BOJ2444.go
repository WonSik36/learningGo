/*
	baekjoon online judge
	problem number 2444
	https://www.acmicpc.net/problem/2444
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

	for i := 1; i <= N-1; i++ {
		for j := N - i; j > 0; j-- {
			w.WriteByte(' ')
		}

		for j := 0; j < 2*i-1; j++ {
			w.WriteByte('*')
		}
		w.WriteByte('\n')
	}

	for i := N; i > 0; i-- {
		for j := 0; j < N-i; j++ {
			w.WriteByte(' ')
		}

		for j := 0; j < 2*i-1; j++ {
			w.WriteByte('*')
		}
		w.WriteByte('\n')
	}

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
