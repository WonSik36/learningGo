/*
	baekjoon online judge
	problem number 16500
	https://www.acmicpc.net/problem/16500

	BackTracking
	high reference:
	https://huiung.tistory.com/117

	Need memoization to remove wrong way
*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var memo []bool

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	target := []byte(getString(r))
	N := getInt(r)
	memo = make([]bool,len(target)+1,len(target)+1)
	for i:=0;i<len(memo);i++ {
		memo[i] = true
	}

	var arr [][]byte
	for i := 0; i < N; i++ {
		arr = append(arr, []byte(getString(r)))
	}

	if dp(0, arr, target) {
		w.WriteString("1\n")
	} else {
		w.WriteString("0\n")
	}

	w.Flush()
}

func dp(pos int, arr [][]byte, target []byte) bool {
	if pos == len(target) {
		return true
	}

	if !memo[pos] {
		return false;
	}

	var res bool = false

	for i := 0; i < len(arr); i++ {
		var flag bool = true
		for j := 0; j < len(arr[i]); j++ {
			// not equal or out of index
			if pos+j >= len(target) || arr[i][j] != target[pos+j] {
				flag = false
				break
			}
		}

		if flag {
			if dp(pos+len(arr[i]), arr, target) {
				res = true
				break
			}
		}
	}

	memo[pos] = res
	return res
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
