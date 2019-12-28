/*
    baekjoon online judge
    problem number 11726
	https://www.acmicpc.net/problem/11726

	Dynamic Programming
	dp[i] = dp[i-1] + dp[i-2]
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

	const mod int = 10007
	N, _ := strconv.Atoi(getString(r))
	var dp []int = make([]int, N+2, N+2)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2

	for i := 3; i <= N; i++ {
		dp[i] = (dp[i-1] + dp[i-2]) % mod
	}

	w.WriteString(strconv.Itoa(dp[N]))
	w.Flush()
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
