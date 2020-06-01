/*
	baekjoon online judge
	problem number 10996
	https://www.acmicpc.net/problem/10996
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	N := getInt(r)

	for i := 0; i < 2*N; i++ {
		if i%2 == 0 {
			for j := 0; j < N; j++ {
				if j%2 == 0 {
					w.WriteByte('*')
				} else {
					w.WriteByte(' ')
				}
			}
		} else {
			for j := 0; j < N; j++ {
				if j%2 == 0 {
					w.WriteByte(' ')
				} else {
					w.WriteByte('*')
				}
			}
		}
		w.WriteByte('\n')
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

func printArr(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Print("\n")
}
