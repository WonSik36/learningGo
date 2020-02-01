/*
	baekjoon online judge
	problem number 1759
	https://www.acmicpc.net/problem/1759

	BackTracking Problem
	reference to problem 15649
	https://yabmoons.tistory.com/115
*/

package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
	"strings"
)

var L int
var C int
var words []byte

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	strTkn := getStrTkn(r)
	L, _ = strconv.Atoi(strTkn[0])
	C, _ = strconv.Atoi(strTkn[1])

	words = getByteTkn(r)
	sort.Slice(words, func(i, j int) bool {
		return words[i] < words[j]
	})

	visited := make([]bool, C, C)

	for i := 0; i <= C-L; i++ {
		visited[i] = true
		dfs(i, 1, visited, w)
		visited[i] = false
	}

	w.Flush()
}

func dfs(cur int, cnt int, visited []bool, w *bufio.Writer) {
	if cnt == L {
		pw := makePassword(visited)
		if isValidPW(pw) {
			w.WriteString(pw)
			w.WriteString("\n")
		}

		return
	}

	for i := cur + 1; i <= C-L+cnt; i++ {
		visited[i] = true
		dfs(i, cnt+1, visited, w)
		visited[i] = false
	}
}

func makePassword(visited []bool) string {
	var sb strings.Builder

	for i := 0; i < len(visited); i++ {
		if visited[i] {
			sb.WriteString(string(words[i]))
		}
	}

	return sb.String()
}

// check the password has more than 1 vowel and more than 2 consonants
func isValidPW(pw string) bool {
	var vowel, cons int = 0, 0

	for i := 0; i < len(pw); i++ {
		if pw[i] == 'a' || pw[i] == 'e' || pw[i] == 'i' || pw[i] == 'o' || pw[i] == 'u' {
			vowel++
		} else {
			cons++
		}
	}

	if vowel >= 1 && cons >= 2 {
		return true
	} else {
		return false
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

func getByteTkn(r *bufio.Reader) []byte {
	strTkn := getStrTkn(r)
	length := len(strTkn)
	res := make([]byte, length, length)

	for i := 0; i < length; i++ {
		res[i] = strTkn[i][0]
	}

	return res
}

func getIntArr(r *bufio.Reader) []int {
	strTkn := getStrTkn(r)
	arr := make([]int, len(strTkn), len(strTkn))

	for i := 0; i < len(arr); i++ {
		arr[i], _ = strconv.Atoi(strTkn[i])
	}

	return arr
}

func printArrByte(arr []byte, w *bufio.Writer) {
	for i := 0; i < len(arr); i++ {
		w.WriteString(string(arr[i]))
		w.WriteString(" ")
	}
	w.WriteString("\n")
}

func printArrBool(arr []bool, w *bufio.Writer) {
	for i := 0; i < len(arr); i++ {
		if arr[i] {
			w.WriteString("#t ")
		} else {
			w.WriteString("#f ")
		}
	}
	w.WriteString("\n")
}
