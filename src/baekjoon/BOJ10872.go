/*
	baekjoon online judge
	problem number 10872
	https://www.acmicpc.net/problem/10872

	recursion vs tail recursion
	https://blog.usejournal.com/tail-recursion-in-go-fb5cf69a0f26
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

	num := getInt(r)
	res := factorialTail(num, 1)

	w.WriteString(strconv.Itoa(res)+"\n")
	w.Flush()
}

// tail recurstion
func factorialTail(n int, res int) int {
	if n==0{
		return res
	}

	return factorialTail(n-1, res*n)
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
