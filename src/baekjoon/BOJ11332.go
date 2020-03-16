/*
	baekjoon online judge
	problem number 11332
	https://www.acmicpc.net/problem/11332

	Online Judge can calculate 10^8 in 1 second
*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type calculate func()

const mod int64 = 100000000

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	N := getInt(r)
	for i := 0; i < N; i++ {
		strTkn := getStrTkn(r)

		max, _ := strconv.ParseInt(strTkn[1], 10, 64)
		cnt, _ := strconv.ParseInt(strTkn[2], 10, 64)
		limit, _ := strconv.ParseInt(strTkn[3], 10, 64)

		var flag bool = true
		switch {
		case strTkn[0] == "O(N)":
			if max*cnt > mod*limit {
				flag = false
			}
		case strTkn[0] == "O(N^2)":
			if max*max*cnt > mod*limit {
				flag = false
			}
		case strTkn[0] == "O(N^3)":
			if max*max*max > mod*limit/cnt {
				flag = false
			}
		case strTkn[0] == "O(2^N)":
			var time int64 = 1
			for i := 1; int64(i) <= max; i++ {
				time *= 2
				if time > mod*limit/cnt {
					flag = false
					break
				}
			}

		case strTkn[0] == "O(N!)":
			var time int64 = 1
			for i := 1; int64(i) <= max; i++ {
				time *= int64(i)
				if time > mod*limit/cnt {
					flag = false
					break
				}
			}
		}

		if flag {
			w.WriteString("May Pass.\n")
		} else {
			w.WriteString("TLE!\n")
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
