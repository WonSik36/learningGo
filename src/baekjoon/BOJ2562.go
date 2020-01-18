/*
	baekjoon online judge
	problem number 2562
	https://www.acmicpc.net/problem/2562
*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const length int = 9

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	
	max := 0
	maxIdx := 1
	for i:=1;i<=length;i++ {
		input := getInt(r)
		if input > max {
			max = input
			maxIdx = i
		}
	}

	var sb strings.Builder
	sb.WriteString(strconv.Itoa(max))
	sb.WriteString("\n")
	sb.WriteString(strconv.Itoa(maxIdx))
	sb.WriteString("\n")

	w.WriteString(sb.String())
	w.Flush()
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
