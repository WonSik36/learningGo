/*
	baekjoon online judge
	problem number 1328
	https://www.acmicpc.net/problem/1328

	Dynamic Programming
	It was very hard to get recurrence relation
	high reference:
	https://jaimemin.tistory.com/536

	1. Recurrence relation !!!
	2. Initialize is important too. 
		I firstly initialize memo value to 0 and if it is 0 than it should be calculated in dp.
		But 0 can be calculated value too.
*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	// "fmt"
)

var memo [101][101][101]int64

const mod int64 = 1000000007

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	arr := getIntArr(r)
	N := arr[0]
	L := arr[1]
	R := arr[2]

	initMemo()
	
	res := dp(N, L, R)
	// printArr(N,memo,w)
	
	// w.WriteString(strconv.Itoa(countZero())+"\n")
	w.WriteString(strconv.FormatInt(res, 10) + "\n")
	w.Flush()
}

func countZero() int {
	cnt := 0
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100; j++ {
			for k := 0; k <= 100; k++ {
				if memo[i][j][k] == 0 {
					cnt++
				}
			}
		}
	}

	return cnt
}

func initMemo() {
	for i := 0; i <= 100; i++ {
		for j := 0; j <= 100; j++ {
			for k := 0; k <= 100; k++ {
				memo[i][j][k] = -1
			}
		}
	}
}

func dp(n int, l int, r int) int64 {
	// fmt.Printf("N: %d, L: %d, R: %d\n",n,l,r)
	if (l == 1 && r == n) || (l == n && r == 1) {
		return 1
	}

	if n == 0 || l == 0 || r == 0 {
		return 0
	}

	if memo[n][l][r] != -1 {
		return memo[n][l][r]
	}

	memo[n][l][r] = (dp(n-1, l, r)*int64(n-2) + dp(n-1, l-1, r) + dp(n-1, l, r-1)) % mod

	return memo[n][l][r]
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

func printArr(len int, arr [101][101][101]int64, w *bufio.Writer) {
	for i := 1; i <= len; i++ {
		for j := 1; j <= len; j++ {
			for k := 1; k <= len; k++ {
				w.WriteString(strconv.FormatInt(arr[i][j][k],10) + " ")
			}
			w.WriteString("\n")
		}
		w.WriteString("\n")
	}
}
