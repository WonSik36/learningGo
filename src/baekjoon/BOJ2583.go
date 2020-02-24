/*
	baekjoon online judge
	problem number 2583
	https://www.acmicpc.net/problem/2583

	DFS Problem
*/

package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

var M int
var N int
var K int
var visited [][]bool
var disabled [][]bool
var res []int
var count int = 0

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	input := getIntArr(r)
	M = input[0]
	N = input[1]
	K = input[2]

	for i := 0; i < M; i++ {
		visited = append(visited, make([]bool, N, N))
		disabled = append(disabled, make([]bool, N, N))
	}

	for i := 0; i < K; i++ {
		input = getIntArr(r)
		markDisabled(input)
	}

	// printArr(disabled, w)

	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			reg := dfs(i, j)

			if reg > 0 {
				res = append(res, reg)
				count++
			}
		}
	}

	w.WriteString(strconv.Itoa(count))
	w.WriteString("\n")

	result := sort.IntSlice(res)
	result.Sort()

	for i := 0; i < len(result); i++ {
		w.WriteString(strconv.Itoa(result[i]))
		w.WriteString(" ")
	}
	w.WriteString("\n")
	w.Flush()
}

func dfs(y int, x int) int {
	if disabled[y][x] {
		return 0
	}

	if visited[y][x] {
		return 0
	}
	visited[y][x] = true

	res := 1

	if x > 0 {
		res += dfs(y, x-1)
	}
	if x < N-1 {
		res += dfs(y, x+1)
	}
	if y > 0 {
		res += dfs(y-1, x)
	}
	if y < M-1 {
		res += dfs(y+1, x)
	}

	return res
}

func markDisabled(input []int) {
	for i := input[1]; i < input[3]; i++ {
		for j := input[0]; j < input[2]; j++ {
			disabled[i][j] = true
		}
	}
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

func printArr(arr [][]bool, w *bufio.Writer) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			if arr[i][j] {
				w.WriteString("1 ")
			} else {
				w.WriteString("0 ")
			}
		}
		w.WriteString("\n")
	}
}
