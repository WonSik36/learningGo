/*
	baekjoon online judge
	problem number 1085
	https://www.acmicpc.net/problem/1085
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
	x, _ := strconv.Atoi(strTkn[0])
	y, _ := strconv.Atoi(strTkn[1])
	width, _ := strconv.Atoi(strTkn[2])
	height, _ := strconv.Atoi(strTkn[3])

	min := 10000
	min = Min(min, x)
	min = Min(min, y)
	min = Min(min, width-x)
	min = Min(min, height-y)

	w.WriteString(strconv.Itoa(min))
	w.WriteString("\n")
	w.Flush()
}

func Min(a int, b int) int {
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
