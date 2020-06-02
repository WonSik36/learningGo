/*
	baekjoon online judge
	problem number 14620
	https://www.acmicpc.net/problem/14620

	Brute Force
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var N int
var COST int = 10000
var arr [][]int
var visited [][]bool

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	N = getInt(r)
	for i := 0; i < N; i++ {
		row := make([]int, N, N)
		arr = append(arr, row)
	}
	for i := 0; i < N; i++ {
		row := make([]bool, N, N)
		visited = append(visited, row)
	}

	for i := 0; i < N; i++ {
		row := getIntArr(r)

		for j := 0; j < N; j++ {
			arr[i][j] = row[j]
		}
	}

	// printArr(arr)

	for i := 1; i < N-1; i++ {
		for j := 1; j < N-1; j++ {
			visit(i, j)
			backtrack(i, j, 1, getCost(i, j))
			unVisit(i, j)
		}
	}

	w.WriteString(strconv.Itoa(COST) + "\n")

	w.Flush()
}

/*
	param: y,x -> pos
	param: cnt -> level of current stage
	param: cost -> cost of current stage
	param: arr [][]int -> map of cost
	param: visited [][]bool -> map of visited before

	return: COST
*/
func backtrack(y int, x int, cnt int, cost int) {
	// fmt.Printf("[%d,%d]:%d cnt: %d\n", y,x,cost,cnt)
	if cnt == 3 {
		COST = Min(COST, cost)
		return
	}

	for i := 1; i < N-1; i++ {
		for j := 1; j < N-1; j++ {
			if !isPossible(i, j) {
				continue
			}

			visit(i, j)
			backtrack(i, j, cnt+1, getCost(i, j)+cost)
			unVisit(i, j)
		}
	}
}

func getCost(y int, x int) int {
	return arr[y-1][x] + arr[y][x-1] + arr[y][x] + arr[y][x+1] + arr[y+1][x]
}

func visit(y int, x int) {
	visited[y-1][x] = true
	visited[y][x-1] = true
	visited[y][x] = true
	visited[y][x+1] = true
	visited[y+1][x] = true
}

func unVisit(y int, x int) {
	visited[y-1][x] = false
	visited[y][x-1] = false
	visited[y][x] = false
	visited[y][x+1] = false
	visited[y+1][x] = false
}

func isPossible(y int, x int) bool {
	if visited[y-1][x] || visited[y][x-1] || visited[y][x] || visited[y][x+1] || visited[y+1][x] {
		return false
	} else {
		return true
	}
}

func Min(a, b int) int {
	if a < b {
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

func printArr(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Printf("%d ", arr[i][j])
		}
		fmt.Println()
	}
	fmt.Print("\n")
}
