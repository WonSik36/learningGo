/*
	baekjoon online judge
	problem number 6246
	https://www.acmicpc.net/problem/6246
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

	arr := getIntArr(r)
	N := arr[0]
	Q := arr[1]

	visited := make([]bool, N+1, N+1)

	for i:=0; i<Q; i++ {
		arr = getIntArr(r)

		start := arr[0]
		gap := arr[1]

		for j:=start;j<=N;j+=gap {
			visited[j] = true
		}
	}

	result := 0
	for i:=1;i<=N;i++ {
		if !visited[i] {
			result++
		}
	}

	w.WriteString(strconv.Itoa(result)+"\n")
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
