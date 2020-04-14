/*
	baekjoon online judge
	problem number 9933
	https://www.acmicpc.net/problem/9933
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

	N := getInt(r)
	arr := make([]string, N, N)

	for i := 0; i < N; i++ {
		arr[i] = getString(r)
	}

	flag := false
	for i := 0; i < N; i++ {
		if flag {
			break
		}

		for j := i; j < N; j++ {
			if inverseEquals(arr[i], arr[j]) {
				resultByte := []byte(arr[i])
				result := resultByte[len(resultByte)/2]

				w.WriteString(strconv.Itoa(len(arr[i])) + " ")
				w.WriteByte(result)
				w.WriteByte('\n')
				flag = true
				break
			}
		}
	}

	w.Flush()
}

func inverseEquals(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	byte1 := []byte(str1)
	byte2 := []byte(str2)

	for i := 0; i < len(byte1); i++ {
		if byte1[i] != byte2[len(byte2)-1-i] {
			return false
		}
	}

	return true
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
