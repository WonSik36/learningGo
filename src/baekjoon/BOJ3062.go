/*
	baekjoon online judge
	problem number 3062
	https://www.acmicpc.net/problem/3062
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

	N := getInt(r)
	for i := 0; i < N; i++ {
		num := getInt(r)
		rNum := getReverseNum(num)
		sum := num + rNum

		if isSymmetry(sum) {
			w.WriteString("YES\n")
		} else {
			w.WriteString("NO\n")
		}
	}

	w.Flush()
}

func isSymmetry(num int) bool {
	var bytes []byte = []byte(strconv.Itoa(num))
	for i := 0; i < len(bytes)/2; i++ {
		if bytes[i] != bytes[len(bytes)-1-i] {
			return false
		}
	}

	return true
}

func getReverseNum(num int) int {
	var bytes []byte = []byte(strconv.Itoa(num))
	var tmp byte

	for i := 0; i < len(bytes)/2; i++ {
		tmp = bytes[i]
		bytes[i] = bytes[len(bytes)-1-i]
		bytes[len(bytes)-1-i] = tmp
	}

	num, _ = strconv.Atoi(string(bytes))
	return num
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
