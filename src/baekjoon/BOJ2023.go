/*
	baekjoon online judge
	problem number 2023
	https://www.acmicpc.net/problem/2023

	Prime Number Problem
	BackTracking Problem
*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var isPrime []bool
var sb strings.Builder
var N int

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	N = getInt(r)
	initPrimeArr(N)
	// printArr(isPrime,w)

	backtrack(0,0)

	w.WriteString(sb.String())
	w.Flush()
}

func backtrack(num int, digit int){
	if digit == N {
		if isPrime[num] {
			sb.WriteString(strconv.Itoa(num))
			sb.WriteString("\n")
		}
		return
	}

	if !isPrime[num] {
		return
	}

	if digit == 0 {
		for i:=1;i<10;i++ {
			backtrack(num*10+i, digit+1)
		}
	}else{
		for i:=1;i<10;i+=2 {
			backtrack(num*10+i, digit+1)
		}
	}

}

func initPrimeArr(N int){
	var Max int = 1
	for i:=0;i<N;i++ {
		Max *= 10
	}

	isPrime = make([]bool, Max, Max)
	for i:=0;i<len(isPrime);i++ {
		isPrime[i] = true
	}

	isPrime[1] = false

	for i:=2;i*i<Max;i++ {
		if isPrime[i] {
			for j:=2*i;j<Max;j+=i {
				isPrime[j] = false
			}
		}
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

func getIntArr(r *bufio.Reader) []int {
	strTkn := getStrTkn(r)
	arr := make([]int, len(strTkn), len(strTkn))

	for i := 0; i < len(arr); i++ {
		arr[i], _ = strconv.Atoi(strTkn[i])
	}

	return arr
}

func printArr(arr []bool, w *bufio.Writer) {
	for i := 0; i < len(arr); i++ {
		if arr[i] {
			w.WriteString(strconv.Itoa(i)+": true\n")
		}else{
			w.WriteString(strconv.Itoa(i)+": false\n")
		}
	}
}
