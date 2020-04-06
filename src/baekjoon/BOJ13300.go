/*
	baekjoon online judge
	problem number 13300
	https://www.acmicpc.net/problem/13300
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var N int
var K int
var input [][]int

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	arr := getIntArr(r)
	N = arr[0]
	K = arr[1]

	// make input arr
	for i := 0; i < 6; i++ {
		tmp := make([]int, 2, 2)
		input = append(input, tmp)
	}

	for i := 0; i < N; i++ {
		arr = getIntArr(r)
		input[arr[1]-1][arr[0]]++
	}

	var sum int = 0
	// find maximum value
	for i := 0; i < 6; i++ {
		for j := 0; j < 2; j++ {
			sum += int(math.Ceil(float64(input[i][j]) / float64(K)))
		}
	}

	w.WriteString(strconv.Itoa(sum) + "\n")
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

func printArr(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("Grade %d\n", i+1)
		for j := 0; j < 2; j++ {
			if j == 0 {
				fmt.Printf("Women: %d\n", arr[i][j])
			} else {
				fmt.Printf("Men: %d\n", arr[i][j])
			}
		}
	}
}
