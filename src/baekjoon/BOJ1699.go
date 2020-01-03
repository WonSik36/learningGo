/*
    baekjoon online judge
    problem number 1699
	https://www.acmicpc.net/problem/1699
	http://blog.naver.com/occidere/220792326120

	Dynamic Programming
*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	N, _ := strconv.Atoi(getString(r))
	var dp []int = make([]int, N+1, N+1)

	// initialize
	for i := 0; i <= N; i++ {
		dp[i] = i
	}

	// set square number's value to 1
	for i := 1; i*i <= N; i++ {
		dp[i*i] = 1
	}

	// if N is square number
	if dp[N] != N {
		w.WriteString(strconv.Itoa(dp[N]) + "\n")
		w.Flush()
		return
	}

	for i := 2; i <= N; i++ {
		for j := 2; j*j <= i; j++ {
			dp[i] = Min(dp[i], dp[i-j*j]+1)
		}
	}

	w.WriteString(strconv.Itoa(dp[N]) + "\n")
	w.Flush()
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
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

func printArr(arr []int, w *bufio.Writer) {
	for i := 0; i < len(arr); i++ {
		w.WriteString(strconv.Itoa(arr[i]) + "\n")
	}
}
