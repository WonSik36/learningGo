/*
	baekjoon online judge
	problem number 5543
	https://www.acmicpc.net/problem/5543
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
	arr := make([]int, 5, 5)

	for i := 0; i < 5; i++ {
		arr[i] = getInt(r)
	}

	hamburger := 10000
	for i := 0; i < 3; i++ {
		hamburger = Min(hamburger, arr[i])
	}

	soda := 10000
	for i := 3; i < 5; i++ {
		soda = Min(soda, arr[i])
	}

	res := hamburger + soda - 50

	w.WriteString(strconv.Itoa(res) + "\n")
	w.Flush()
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

func printArr(arr []int, w *bufio.Writer) {
	for i := 0; i < len(arr); i++ {
		w.WriteString(strconv.Itoa(arr[i]) + "\n")
	}
}
