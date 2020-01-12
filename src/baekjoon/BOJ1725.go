/*
   	baekjoon online judge
   	problem number 1725
   	https://www.acmicpc.net/problem/1725
	https://www.acmicpc.net/blog/view/12

	deep reference

   	stack problem
   	Histogram
   	same with 6549
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
	var res strings.Builder

	N := getInt(r)
	var arr []int = make([]int, N, N)
	for i := 0; i < N; i++ {
		arr[i] = getInt(r)
	}

	st := MakeStack()
	var max int = 0

	for i := 0; i < N; i++ {
		for !st.Empty() && arr[st.Top()] > arr[i] {
			var h int = arr[st.Pop()]
			var w int

			if st.Empty() {
				w = i
			} else {
				w = i - st.Top() - 1
			}

			max = Max(max, w*h)
		}

		st.Push(i)
	}

	for !st.Empty() {
		var h int = arr[st.Pop()]
		var w int

		if st.Empty() {
			w = N
		} else {
			w = N - st.Top() - 1
		}

		max = Max(max, w*h)
	}

	res.WriteString(strconv.Itoa(max))
	res.WriteString("\n")

	w.WriteString(res.String())
	w.Flush()
}

func Max(a int, b int) int {
	if a > b {
		return a
	} else {
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

func printArr(arr []int, w *bufio.Writer) {
	for i := 0; i < len(arr); i++ {
		w.WriteString(strconv.Itoa(arr[i]) + "\n")
	}
}

/*************** Stack and Node ***************/

type Stack struct {
	size int
	head *Node
}

func MakeStack() *Stack {
	return &Stack{0, nil}
}

func (st *Stack) Push(value int) {
	var node *Node
	if st.size == 0 {
		node = MakeNode(value, nil)
	} else {
		node = MakeNode(value, st.head)
	}
	st.head = node
	st.size++
}

func (st *Stack) Pop() int {
	if st.size == 0 {
		panic("There are no element in stack")
	}

	v := st.head.value
	st.head = st.head.next
	st.size--
	return v
}

func (st *Stack) Size() int {
	return st.size
}

func (st *Stack) Empty() bool {
	if st.size == 0 {
		return true
	} else {
		return false
	}
}

func (st *Stack) Top() int {
	if st.size == 0 {
		panic("There are no element in stack")
	} else {
		return st.head.value
	}
}

type Node struct {
	value int
	next  *Node
}

func MakeNode(value int, next *Node) *Node {
	return &Node{value, next}
}
