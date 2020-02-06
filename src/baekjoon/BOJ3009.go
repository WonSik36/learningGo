/*
	baekjoon online judge
	problem number 3009
	https://www.acmicpc.net/problem/3009
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

	point1 := getIntArr(r)
	point2 := getIntArr(r)
	point3 := getIntArr(r)

	var x, y int

	if point1[0] == point2[0] {
		x = point3[0]
	} else if point1[0] == point3[0] {
		x = point2[0]
	} else { // point2[0] == point3[0]
		x = point1[0]
	}

	if point1[1] == point2[1] {
		y = point3[1]
	} else if point1[1] == point3[1] {
		y = point2[1]
	} else { // point2[1] == point3[1]
		y = point1[1]
	}

	w.WriteString(strconv.Itoa(x))
	w.WriteString(" ")
	w.WriteString(strconv.Itoa(y))
	w.WriteString("\n")

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
