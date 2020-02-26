/*
	baekjoon online judge
	problem number 2468
	https://www.acmicpc.net/problem/2468

	light reference:
	https://www.acmicpc.net/board/view/27135
*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var N int
var maps [][]int
var visited [][]bool

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	N = getInt(r)
	for i := 0; i < N; i++ {
		maps = append(maps, make([]int, N, N))
	}

	min, max := 1000, 0
	for i := 0; i < N; i++ {
		arr := getIntArr(r)
		for j := 0; j < N; j++ {
			maps[i][j] = arr[j]
			min = Min(min, maps[i][j])
			max = Max(max, maps[i][j])
		}
	}

	if min == max {
		w.WriteString("1\n")
	} else {
		var res int = 1
		for i := min; i < max; i++ {
			res = Max(res, dfs(i))
		}
		w.WriteString(strconv.Itoa(res))
		w.WriteString("\n")
	}

	w.Flush()
}

func dfs(height int) int {
	var region int = 0
	// initialize visited
	visited = nil
	for i := 0; i < N; i++ {
		visited = append(visited, make([]bool, N, N))
	}

	for i := 0; i < len(maps); i++ {
		for j := 0; j < len(maps[i]); j++ {
			region += _dfs(i, j, height)
		}
	}

	return region
}

func _dfs(y, x, height int) int {
	if maps[y][x] <= height {
		return 0
	}

	if visited[y][x] {
		return 0
	}
	visited[y][x] = true

	if y > 0 {
		_dfs(y-1, x, height)
	}
	if y < N-1 {
		_dfs(y+1, x, height)
	}
	if x > 0 {
		_dfs(y, x-1, height)
	}
	if x < N-1 {
		_dfs(y, x+1, height)
	}

	return 1
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
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

func printArr(arr []int, w *bufio.Writer) {
	for i := 0; i < len(arr); i++ {
		w.WriteString(strconv.Itoa(arr[i]) + "\n")
	}
}
