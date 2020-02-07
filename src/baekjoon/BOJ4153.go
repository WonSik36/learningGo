/*
	baekjoon online judge
	problem number 4153
	https://www.acmicpc.net/problem/4153
*/

package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var arr []int
	for {
		arr = getIntArr(r)
		sort.IntSlice(arr).Sort()

		// check input value to end
		if arr[0] == 0 && arr[1] == 0 && arr[2] == 0 {
			break
		}

		// two rules
		// 1. a+b > c
		// 2. a^2+b^2 = c^2
		if arr[0]+arr[1] > arr[2] && arr[0]*arr[0]+arr[1]*arr[1] == arr[2]*arr[2] {
			w.WriteString("right\n")
		} else {
			w.WriteString("wrong\n")
		}
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
