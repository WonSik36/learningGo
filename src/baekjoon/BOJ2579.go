/*
	baekjoon online judge
	problem number 2579
	https://www.acmicpc.net/problem/2579

	Dynamic Programming
*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var memo []int

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	N := getInt(r)
	memo = make([]int, N, N)
	arr := make([]int, N, N)
	for i := 0; i < N; i++ {
		arr[i] = getInt(r)
	}

	/* exception case */
	if N == 1 {
		w.WriteString(strconv.Itoa(arr[0]) + "\n")
		w.Flush()
		return
	} else if N == 2 {
		w.WriteString(strconv.Itoa(arr[0]+arr[1]) + "\n")
		w.Flush()
		return
	} else if N == 3 {
		w.WriteString(strconv.Itoa(Max(arr[0]+arr[2], arr[1]+arr[2])) + "\n")
		w.Flush()
		return
	}

	memo[0] = arr[0]
	memo[1] = arr[0] + arr[1]
	memo[2] = Max(arr[0]+arr[2], arr[1]+arr[2])
	for i := 3; i < N; i++ {
		memo[i] = Max(arr[i]+arr[i-1]+memo[i-3], arr[i]+memo[i-2])
	}

	w.WriteString(strconv.Itoa(memo[N-1]) + "\n")
	w.Flush()
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
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

func printArr(arr []int, w *bufio.Writer) {
	for i := 0; i < len(arr); i++ {
		w.WriteString(strconv.Itoa(arr[i]) + "\n")
	}
}
