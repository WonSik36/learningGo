/*
	baekjoon online judge
	problem number 7579
	https://www.acmicpc.net/problem/7579

	Dynamic Programming
	Applicaiton of Knapsack problem
	High Reference: https://huiung.tistory.com/120
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
var M int
var memory []int
var cost []int
var memo [][]int

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	arr := getIntArr(r)
	N = arr[0]
	M = arr[1]

	memory = getIntArr(r)
	cost = getIntArr(r)

	// initialize memo for dp
	for i := 0; i <= N; i++ {
		memo = append(memo, make([]int, 10001, 10001))
	}

	var result int = 1000000
	for i := 1; i <= N; i++ {
		for j := 0; j <= 10000; j++ {
			if cost[i-1] > j {
				memo[i][j] = memo[i-1][j]
				continue
			}

			memo[i][j] = Max(memo[i-1][j], memo[i-1][j-cost[i-1]]+memory[i-1]) // Caution! memory[i-1], cost[i-1] is for indexing

			if memo[i][j] >= M {
				// fmt.Printf("result:%d, j:%d\n",result,memo[i][j])
				result = Min(result, j)
			}
		}
	}

	w.WriteString(strconv.Itoa(result) + "\n")
	w.Flush()
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
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

func printArr(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Print("\n")
}
