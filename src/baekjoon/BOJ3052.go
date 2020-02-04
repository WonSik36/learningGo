/*
	baekjoon online judge
	problem number 3052
	https://www.acmicpc.net/problem/3052
*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const mod int = 42

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	// get input array
	arr := []int{}
	for i := 0; i < 10; i++ {
		arr = append(arr, getInt(r))
	}

	// modulo operation
	for i := 0; i < len(arr); i++ {
		arr[i] %= mod
	}

	// get set of input array
	remain := []int{}
	for i := 0; i < len(arr); i++ {
		if !contains(remain, arr[i]) {
			remain = append(remain, arr[i])
		}
	}

	w.WriteString(strconv.Itoa(len(remain)))
	w.WriteString("\n")
	w.Flush()
}

func contains(arr []int, num int) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == num {
			return true
		}
	}
	return false
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
