/*
	baekjoon online judge
	problem number 11729
	https://www.acmicpc.net/problem/11729

	Tower of Hanoi Problem(Recursion)
*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var cnt int = 0

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	var sb strings.Builder

	N := getInt(r)
	hanoi(1, 2, 3, N, &sb)

	w.WriteString(strconv.Itoa(cnt) + "\n")
	w.WriteString(sb.String())
	w.Flush()
}

func hanoi(src int, pass int, dst int, height int, sb *strings.Builder) {
	if height <= 0 {
		return
	}

	// move 1~height-1 plates from src to pass by passing through dst
	hanoi(src, dst, pass, height-1, sb)

	// move height plate from src to dst
	sb.WriteString(strconv.Itoa(src))
	sb.WriteString(" ")
	sb.WriteString(strconv.Itoa(dst))
	sb.WriteString("\n")
	cnt++

	// move 1~height-1 from pass to dst by passing through src
	hanoi(pass, src, dst, height-1, sb)
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
