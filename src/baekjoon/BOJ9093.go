/*
	baekjoon online judge
	problem number 9093
	https://www.acmicpc.net/problem/9093
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
	var sb strings.Builder

	N := getInt(r)
	for i := 0; i < N; i++ {
		sb.Reset()
		strTkn := getStrTkn(r)
		for j := 0; j < len(strTkn); j++ {
			tkn := []byte(strTkn[j])
			for k := len(tkn) - 1; k >= 0; k-- {
				sb.WriteByte(tkn[k])
			}

			if j != len(strTkn)-1 {
				sb.WriteByte(' ')
			} else {
				sb.WriteByte('\n')
			}
		}

		w.WriteString(sb.String())
	}

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
