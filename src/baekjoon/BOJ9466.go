/*
	baekjoon online judge
	problem number 9466
	https://www.acmicpc.net/problem/9466

	DFS Problem, finding node which is not in cycle
	Deep Reference:
	https://mygumi.tistory.com/107
	https://github.com/hotehrud/acmicpc/blob/master/algorithm/graph/9466.java

	The make function spends time quite a bit of time
	The difference between deprecated and reference is how much they call the make function.
*/

package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	// "fmt"
)

var next []int
var srcNode []int
var visitOrder []int

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	N := getInt(r)
	for i := 0; i < N; i++ {
		M := getInt(r)
		next = make([]int, 1, 1)
		arr := getIntArr(r)
		next = append(next, arr...)
		srcNode = make([]int, M+1, M+1)
		visitOrder = make([]int, M+1, M+1)

		// printArr(next,w)
		res := dfs(M)
		w.WriteString(strconv.Itoa(res) + "\n")
	}

	w.Flush()
}

func dfs(N int) int {
	mem := 0
	for i := 1; i <= N; i++ {
		if visitOrder[i] == 0 {
			mem += _dfs(i, i)
		}
	}

	return N - mem
}

func _dfs(src, cur int) int {
	cnt := 1

	for {
		if visitOrder[cur] != 0 {
			if srcNode[cur] != src {
				return 0
			}
			return cnt - visitOrder[cur]
		}

		srcNode[cur] = src
		visitOrder[cur] = cnt
		cur = next[cur]

		cnt++
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
		w.WriteString(strconv.Itoa(arr[i]) + " ")
	}
	w.WriteString("\n")
}

/********** Deprecated *********/
/*
var next []int
var done []bool
var visited []bool
var visitOrder []int

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	N := getInt(r)
	for i := 0; i < N; i++ {
		M := getInt(r)
		next = make([]int, 1, 1)
		arr := getIntArr(r)
		next = append(next, arr...)

		// printArr(next,w)
		res := dfs(M)
		w.WriteString(strconv.Itoa(res) + "\n")
	}

	w.Flush()
}

func dfs(N int) int {
	done = make([]bool, N+1, N+1)
	cntCycleMember := 0

	for i := 1; i <= N; i++ {
		if done[i] {
			continue
		}
		visited = make([]bool, N+1, N+1)
		visitOrder = make([]int, N+1, N+1)

		// fmt.Printf("_dfs %d start\n",i)
		cnt := _dfs(i, i, 0)
		cntCycleMember += cnt
		// fmt.Printf("_dfs %d end, cnt: %d\n",i,cnt)
	}

	return N - cntCycleMember
}

func _dfs(src, cur, cnt int) int {
	// fmt.Printf("src:%d, cur:%d, cnt:%d\n",src,cur,cnt)
	if visited[cur] {
		return cnt - visitOrder[cur]
	}
	// visited by previous dfs
	if done[cur] {
		return 0
	}

	done[cur] = true
	visited[cur] = true
	visitOrder[cur] = cnt

	return _dfs(src, next[cur], cnt+1)
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
		w.WriteString(strconv.Itoa(arr[i]) + " ")
	}
	w.WriteString("\n")
}
*/