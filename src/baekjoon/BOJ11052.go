/*
    baekjoon online judge
    problem number 11052
	https://www.acmicpc.net/problem/11052
	https://yabmoons.tistory.com/23

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
	var arr []int = make([]int, N+1, N+1)
	var dp []int = make([]int, N+1, N+1)

	strTkn := getStrTkn(r)
	for i := 1; i <= N; i++ {
		arr[i], _ = strconv.Atoi(strTkn[i-1])
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= i; j++ {
			dp[i] = Max(dp[i], dp[i-j]+arr[j])
		}
	}

	w.WriteString(strconv.Itoa(dp[N]) + "\n")
	w.Flush()
}

func Max(a, b int) int {
	if a > b {
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
