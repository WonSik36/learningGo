/*
    baekjoon online judge
    problem number 10828
	https://www.acmicpc.net/problem/10828

	make stack

	introduce strings.Builder
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
	var sb strings.Builder

	N := getInt(r)
	st := MakeStack()

	for i := 0; i < N; i++ {
		sb.Reset()
		strTkn := getStrTkn(r)

		switch {
		case strTkn[0] == "push":
			v, _ := strconv.Atoi(strTkn[1])
			st.Push(v)
		case strTkn[0] == "pop":
			sb.WriteString(strconv.Itoa(st.Pop()))
			sb.WriteString("\n")
			w.WriteString(sb.String())
		case strTkn[0] == "size":
			sb.WriteString(strconv.Itoa(st.Size()))
			sb.WriteString("\n")
			w.WriteString(sb.String())
		case strTkn[0] == "empty":
			if st.Empty() {
				w.WriteString("1\n")
			} else {
				w.WriteString("0\n")
			}
		case strTkn[0] == "top":
			sb.WriteString(strconv.Itoa(st.Top()))
			sb.WriteString("\n")
			w.WriteString(sb.String())
		}

	}

	w.Flush()
}

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
		return -1
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
		return -1
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
