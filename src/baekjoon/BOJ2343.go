/*
	baekjoon online judge
	problem number 2343
	https://www.acmicpc.net/problem/2343

	binary search with parametric search
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

	strTkn := getStrTkn(r)
	N,_ := strconv.Atoi(strTkn[0])
	M,_ := strconv.Atoi(strTkn[1])

	max := 0	// max value in array
	sum := 0	// sum of value in array
	strTkn = getStrTkn(r)
	var arr []int = make([]int, N, N)
	for i:=0;i<N;i++ {
		v,_ := strconv.Atoi(strTkn[i])
		arr[i] = v
		max = Max(max, v)
		sum += v
	}

	// binary search
	// minimum value of answer will be maximum value of input array
	// maximum value of answer will be sum of input array
	var lo, hi int = max-1, sum+1
	var mid int
	for lo+1<hi {
		mid = (lo+hi)/2
		if isValid(mid, M, arr) {
			hi = mid
		}else{
			lo = mid
		}
	}
	
	w.WriteString(strconv.Itoa(hi)+"\n")
	w.Flush()
}

// check validity of parameter
func isValid(param int, totalNum int, arr []int) bool {
	var storage int = 0
	var num int = 1

	for i:=0;i<len(arr);i++ {
		// if current cd's storage over input param
		if storage+arr[i] > param {
			storage = 0
			num++
			
			if num > totalNum {
				return false
			}
		}

		storage += arr[i]
	}

	if num <= totalNum {
		return true
	}else{
		return false
	}
}

func Max(a int, b int) int {
	if a>b{
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
