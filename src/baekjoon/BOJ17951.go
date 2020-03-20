/*
	baekjoon online judge
	problem number 17951
	https://www.acmicpc.net/problem/17951

	I misunderstood the problem at first.
	Binary Search with Parametirc Search
	reference: https://sys09270883.github.io/ps/6/#17952%EB%B2%88-%EA%B3%BC%EC%A0%9C%EB%8A%94-%EB%81%9D%EB%82%98%EC%A7%80-%EC%95%8A%EC%95%84
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

	arr := getIntArr(r)
	N := arr[0]
	K := arr[1]
	arr = getIntArr(r)

	var lo, hi, mid int = 0, 0, 0
	var sum, cnt int = 0, 0

	for i := 0; i < N; i++ {
		hi += arr[i]
	}
	hi++

	for lo+1 < hi {
		mid = (lo + hi) / 2
		sum = 0
		cnt = 0
		for i := 0; i < N; i++ {
			sum += arr[i]

			if sum >= mid {
				cnt++
				sum = 0
			}
		}

		if cnt >= K {
			lo = mid
		} else {
			hi = mid
		}
	}

	w.WriteString(strconv.Itoa(lo) + "\n")
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
