/*
	baekjoon online judge
	problem number 2512
	https://www.acmicpc.net/problem/2512

	binary search
	Using updated binary search code
	But I missed hi's first value which is max value of input array
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

	getInt(r)	// get N
	max, arr := getIntArr(r)
	limit := getInt(r)

	res := binarySearch(max, arr, limit)

	sb.WriteString(strconv.Itoa(res))
	sb.WriteString("\n")

	w.WriteString(sb.String())
	w.Flush()
}

func binarySearch(max int, arr []int, limit int) int {
	var lo, hi int = 0, max+1

	for lo+1 < hi {
		mid := (lo + hi) / 2
		if isValid(mid, limit, arr) {
			lo = mid
		} else {
			hi = mid
		}
	}

	return lo
}

func isValid(param int, limit int, finance []int) bool {
	var sum int = 0

	for i := 0; i < len(finance); i++ {
		sum += Min(param, finance[i])
	}

	if sum <= limit {
		return true
	} else {
		return false
	}
}

func Min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Max(a int, b int) int {
	if a > b {
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

func getIntArr(r *bufio.Reader) (int,[]int) {
	strTkn := getStrTkn(r)
	arr := make([]int, len(strTkn), len(strTkn))
	max := 0

	for i := 0; i < len(arr); i++ {
		arr[i], _ = strconv.Atoi(strTkn[i])
		max = Max(max,arr[i])
	}

	return max, arr
}

func printArr(arr []int, w *bufio.Writer) {
	for i := 0; i < len(arr); i++ {
		w.WriteString(strconv.Itoa(arr[i]) + "\n")
	}
}
