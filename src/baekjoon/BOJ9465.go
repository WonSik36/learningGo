/*
    baekjoon online judge
    problem number 9465
	https://www.acmicpc.net/problem/9465

	dynamic programming
	dp[0][i] = arr[0][i] + max(dp[1][i-1], dp[1][i-2])
	dp[1][i] = arr[1][i] + max(dp[0][i-1], dp[0][i-2])
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

	for i := 0; i < N; i++ {
		// make 2d array for input
		M, _ := strconv.Atoi(getString(r))
		var arr [2][]int
		for j := 0; j < 2; j++ {
			arr[j] = make([]int, M+1, M+1)
		}

		// get data from input
		for j := 0; j < 2; j++ {
			strTkn := getStrTkn(r)
			for k := 0; k < M; k++ {
				arr[j][k], _ = strconv.Atoi(strTkn[k])
			}
		}

		// printArr(arr,w)
		score := getMaxScore(arr, M)
		w.WriteString(strconv.Itoa(score) + "\n")
	}

	w.Flush()
}

func getMaxScore(arr [2][]int, length int) int {
	var dp [2][]int
	for i := 0; i < 2; i++ {
		dp[i] = make([]int, len(arr[0]), len(arr[0]))
	}

	dp[0][0] = arr[0][0]
	dp[1][0] = arr[1][0]
	dp[0][1] = arr[0][1] + arr[1][0]
	dp[1][1] = arr[1][1] + arr[0][0]

	for i := 2; i < length; i++ {
		dp[0][i] = arr[0][i] + Max(dp[1][i-1], dp[1][i-2])
		dp[1][i] = arr[1][i] + Max(dp[0][i-1], dp[0][i-2])
	}

	return Max(dp[0][length-1], dp[1][length-1])
}

func Max(a int, b int) int {
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

func printArr(arr [2][]int, w *bufio.Writer) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			w.WriteString(strconv.Itoa(arr[i][j]) + " ")
		}
		w.WriteString("\n")
	}
}
