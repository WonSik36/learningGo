/*
    baekjoon online judge
    problem number 2294
	https://www.acmicpc.net/problem/2294

	Dynamic Programming

	C(i): value of ith coin
    DP(k): minimum combination to make k

	DP(0) = 0
	DP(C(i)) = 1
    DP(k) = Min(DP(k), DP(k-C(i)) + DP(C(i)))
*/

package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

const INF int = 1000000

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	strTkn := getStrTkn(r)
	N, _ := strconv.Atoi(strTkn[0])
	K, _ := strconv.Atoi(strTkn[1])
	arr := make([]int, N, N)

	for i := 0; i < N; i++ {
		arr[i], _ = strconv.Atoi(getString(r))
	}

	/* sort array */
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	/* 	set's length is same with arr
	but it doesn't have dupliacted value and left space is filled with 0 */
	set := arr2Set(arr)

	// printArr(set,w)

	dp := make([]int, K+1, K+1)

	/* set dp's value to infinite */
	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		dp[i] = INF
	}

	/* set dp[coin[i]] = 1 */
	for i := 0; i < len(set); i++ {
		if set[i] == 0 {
			break
		}

		if set[i] >= len(dp) {
			continue
		}

		dp[set[i]] = 1
	}

	/* dynamic programming */
	for i := 0; i < len(set); i++ {
		if set[i] == 0 {
			break
		}

		for j := set[i]; j <= K; j++ {
			dp[j] = min(dp[j], dp[j-set[i]]+dp[set[i]])
		}
	}

	if dp[K] == INF {
		w.WriteString("-1\n")
	} else {
		w.WriteString(strconv.Itoa(dp[K]) + "\n")
	}

	w.Flush()
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func arr2Set(arr []int) []int {
	var ret []int = make([]int, len(arr), len(arr))
	j := 0
	ret[j] = arr[0]
	for i := 1; i < len(arr); i++ {
		if ret[j] != arr[i] {
			j++
			ret[j] = arr[i]
		}
	}

	return ret
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
		w.WriteString(strconv.Itoa(arr[i]) + " ")
	}
	w.WriteString("\n")
}
