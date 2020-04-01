/*
	baekjoon online judge
	problem number 12847
	https://www.acmicpc.net/problem/12847
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

	arr := getIntArr(r)
	N := arr[0]
	M := arr[1]
	sum := make([]int64, N+1, N+1)

	arr = getIntArr(r)
	for i := 1; i <= N; i++ {
		sum[i] = sum[i-1] + int64(arr[i-1])
	}

	var max int64 = 0
	for i := 0; i <= N-M; i++ {
		max = Max(max, sum[i+M]-sum[i])
	}

	w.WriteString(strconv.FormatInt(max, 10) + "\n")
	w.Flush()
}

func Max(a, b int64) int64 {
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

func printArr(arr []int64, w *bufio.Writer) {
	for i := 0; i < len(arr); i++ {
		w.WriteString(strconv.FormatInt(arr[i], 10) + "\n")
	}
}
