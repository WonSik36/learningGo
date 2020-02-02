/*
	baekjoon online judge
	problem number 1987
	https://www.acmicpc.net/problem/1987

	BackTracking Problem
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

	strTkn := getStrTkn(r)
	rows, _ := strconv.Atoi(strTkn[0])
	cols, _ := strconv.Atoi(strTkn[1])

	board := [][]byte{}
	for i := 0; i < rows; i++ {
		str := getString(r)
		arr := make([]byte, cols, cols)
		for j := 0; j < cols; j++ {
			arr[j] = str[j]
		}
		board = append(board, arr)
	}

	visited := []byte{}
	visited = append(visited, board[0][0])
	res := dfs(0, 0, 1, board, visited)

	w.WriteString(strconv.Itoa(res))
	w.WriteString("\n")
	w.Flush()
}

func dfs(y int, x int, cnt int, board [][]byte, visited []byte) int {
	max := cnt

	if x < len(board[0])-1 && !contains(board[y][x+1], visited) {
		visited = append(visited, board[y][x+1])
		max = Max(max, dfs(y, x+1, cnt+1, board, visited))
		visited = visited[:len(visited)-1]
	}
	if x > 0 && !contains(board[y][x-1], visited) {
		visited = append(visited, board[y][x-1])
		max = Max(max, dfs(y, x-1, cnt+1, board, visited))
		visited = visited[:len(visited)-1]
	}
	if y < len(board)-1 && !contains(board[y+1][x], visited) {
		visited = append(visited, board[y+1][x])
		max = Max(max, dfs(y+1, x, cnt+1, board, visited))
		visited = visited[:len(visited)-1]
	}
	if y > 0 && !contains(board[y-1][x], visited) {
		visited = append(visited, board[y-1][x])
		max = Max(max, dfs(y-1, x, cnt+1, board, visited))
		visited = visited[:len(visited)-1]
	}

	return max
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func contains(char byte, visited []byte) bool {
	for i := 0; i < len(visited); i++ {
		if visited[i] == char {
			return true
		}
	}
	return false
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

func printArray(board [][]byte, w *bufio.Writer) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			w.WriteString(string(board[i][j]))
			w.WriteString(" ")
		}
		w.WriteString("\n")
	}
}
