/*
    baekjoon online judge
    problem number 2193
	https://www.acmicpc.net/problem/2193

	Dynamic Programming
	dp[i][0] = dp[i-1][0] + dp[i-1][1]
	dp[i][1] = dp[i-1][0]
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
	var dp [][2]int64 = make([][2]int64, N+1, N+1)
	dp[1][0] = 0
	dp[1][1] = 1

	for i := 2; i <= N; i++ {
		dp[i][0] = dp[i-1][0] + dp[i-1][1]
		dp[i][1] = dp[i-1][0]
	}

	res := dp[N][0] + dp[N][1]
	w.WriteString(strconv.FormatInt(res, 10))
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
