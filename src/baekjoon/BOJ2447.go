/*
	baekjoon online judge
	problem number 2447
	https://www.acmicpc.net/problem/2447

	recursion problem
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

	N := getInt(r)
	var arr [][]bool
	for i := 0; i < N; i++ {
		tmp := make([]bool, N, N)
		arr = append(arr, tmp)
	}

	rec(0, 0, N, arr)

	print2DArr(arr, w)
	w.Flush()
}

func rec(y int, x int, scale int, arr [][]bool) {
	if scale == 1 {
		arr[y][x] = true
		return
	}

	rec(y, x, scale/3, arr)
	rec(y, x+scale/3, scale/3, arr)
	rec(y, x+scale/3*2, scale/3, arr)

	rec(y+scale/3, x, scale/3, arr)
	rec(y+scale/3, x+scale/3*2, scale/3, arr)

	rec(y+scale/3*2, x, scale/3, arr)
	rec(y+scale/3*2, x+scale/3, scale/3, arr)
	rec(y+scale/3*2, x+scale/3*2, scale/3, arr)
}

func print2DArr(arr [][]bool, w *bufio.Writer) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			if arr[i][j] {
				w.WriteString("*")
			} else {
				w.WriteString(" ")
			}
		}
		w.WriteString("\n")
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
