/*
	baekjoon online judge
	problem number 8911
	https://www.acmicpc.net/problem/8911
*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

var x = [4]int{0,-1,0,1}
var y = [4]int{1,0,-1,0}

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	N := getInt(r)
	for i:=0;i<N;i++ {
		res := solve(r)
		w.WriteString(strconv.Itoa(res)+"\n")
	}

	w.Flush()
}

func solve(r *bufio.Reader) int {
	var input []byte = []byte(getString(r))
	var curPos = [3]int{0,0,0};
	minX, minY, maxX, maxY := 0,0,0,0

	for i:=0;i<len(input);i++ {
		switch input[i] {
		case 'F':
			curPos[0] += x[curPos[2]]
			curPos[1] += y[curPos[2]]
		case 'B':
			curPos[0] -= x[curPos[2]]
			curPos[1] -= y[curPos[2]]
		case 'L':
			curPos[2] = rotate(curPos[2],1)
		case 'R':
			curPos[2] = rotate(curPos[2],-1)
		}

		minX = Min(minX, curPos[0])
		minY = Min(minY, curPos[1])
		maxX = Max(maxX, curPos[0])
		maxY = Max(maxY, curPos[1])
	}

	return (maxX - minX) * (maxY - minY)
}

func Max(a,b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func Min(a,b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func rotate(src, direction int) int {
	res := src+direction

	if res == 4 {
		return 0
	}

	if res == -1 {
		return 3
	}

	return res
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
