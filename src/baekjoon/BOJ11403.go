/*
	baekjoon online judge
	problem number 11403
	https://www.acmicpc.net/problem/11403

	reference:
	https://www.acmicpc.net/board/view/31656

	Solve by using dfs
*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var visited []bool
var graph [][]bool
var N int

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	N = getInt(r)
	for i := 0; i < N; i++ {
		graph = append(graph, make([]bool, N, N))
	}

	for i := 0; i < N; i++ {
		arr := getIntArr(r)
		for j := 0; j < N; j++ {
			if arr[j] == 1 {
				graph[i][j] = true
			}
		}
	}

	// printArr(graph, w)
	for i := 0; i < N; i++ {
		visited = make([]bool, N, N)
		dfs(i)
		printVisited(visited, w)
	}

	w.Flush()
}

func dfs(idx int) {
	for i := 0; i < N; i++ {
		// there is a way and not visited before
		if graph[idx][i] && !visited[i] {
			visited[i] = true
			dfs(i)
		}
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

func printVisited(arr []bool, w *bufio.Writer) {
	for i := 0; i < len(arr); i++ {
		if arr[i] {
			w.WriteString("1 ")
		} else {
			w.WriteString("0 ")
		}
	}
	w.WriteString("\n")
}
