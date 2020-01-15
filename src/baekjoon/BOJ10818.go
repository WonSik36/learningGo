/*
   	baekjoon online judge
   	problem number 10818
	https://www.acmicpc.net/problem/10818
	   
	array problem
*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

const MAX_NUM = 1000000000
const MIN_NUM = -1000000000

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	getInt(r)	// get N
	arr := getIntArr(r)

	max := MIN_NUM
	min := MAX_NUM

	for i:=0;i<len(arr);i++ {
		max = Max(max, arr[i])
		min = Min(min, arr[i])
	}

	var sb strings.Builder
	sb.WriteString(strconv.Itoa(min))
	sb.WriteString(" ")
	sb.WriteString(strconv.Itoa(max))

	w.WriteString(sb.String())
	w.Flush()
}

func Max(a int, b int) int {
	if a>b{
		return a
	}else{
		return b
	}
}

func Min(a int, b int) int {
	if a<b{
		return a
	}else{
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

func getIntArr(r *bufio.Reader) []int{
	strTkn := getStrTkn(r)
	arr := make([]int, len(strTkn), len(strTkn))

	for i:=0;i<len(arr);i++ {
		arr[i],_ = strconv.Atoi(strTkn[i])
	}

	return arr
}

func printArr(arr []int, w *bufio.Writer) {
	for i := 0; i < len(arr); i++ {
		w.WriteString(strconv.Itoa(arr[i]) + "\n")
	}
}
