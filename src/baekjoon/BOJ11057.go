/*
    baekjoon online judge
    problem number 11057
	https://www.acmicpc.net/problem/11057

	Dynamic Programming
	dp[i][j] = dp[i-1][0] + ... + dp[i-1][j]
	ex) dp[i][5] = dp[i-1][0] + dp[i-1][1] + dp[i-1][2] + dp[i-1][3] + dp[i-1][4] + dp[i-1][5]
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const mod int = 10007

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	N, _ := strconv.Atoi(getString(r))
	var dp [][10]int = make([][10]int, N, N)

	for i := 0; i < 10; i++ {
		dp[0][i] = 1
	}

	// dp[i][j] = dp[i-1][0] + ... + dp[i-1][j]
	// ex) dp[i][5] = dp[i-1][0] + dp[i-1][1] + dp[i-1][2] + dp[i-1][3] + dp[i-1][4] + dp[i-1][5]
	for i := 1; i < N; i++ {
		for j := 0; j < 10; j++ {
			for k := 0; k <= j; k++ {
				dp[i][j] = (dp[i][j] + dp[i-1][k]) % mod
			}
		}
	}

	// printArr(dp)

	sum := 0
	for i := 0; i < 10; i++ {
		sum = (sum + dp[N-1][i]) % mod
	}

	w.WriteString(fmt.Sprintf("%d\n", sum))
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

func printArr(arr [][10]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Println()
	}
}
