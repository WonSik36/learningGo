/*
	baekjoon online judge
	problem number 5624
	https://www.acmicpc.net/problem/5624

	Referece:
	https://jaimemin.tistory.com/876

	Make Triple Loop to Double Loop
	A+B+C = D -> A+B = D-C
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const MaxValue int = 200000  // max value of A+B = 100000+100000 = 200000
const MaxArrIdx int = 400001 // range of A+B: -200000 ~ 200000

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var visited [MaxArrIdx]bool
	result := 0

	N := getInt(r)
	arr := getIntArr(r)

	for i := 0; i < N; i++ {
		// calculate D-C
		for j := 0; j < i; j++ {
			if visited[arr[i]-arr[j]+MaxValue] {
				result++
				break
			}
		}

		// calculate A+B
		for j := 0; j <= i; j++ {
			visited[arr[i]+arr[j]+MaxValue] = true
		}
	}

	w.WriteString(strconv.Itoa(result) + "\n")
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
