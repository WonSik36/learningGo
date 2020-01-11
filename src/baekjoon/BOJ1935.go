/*
   baekjoon online judge
   problem number 1935
   https://www.acmicpc.net/problem/1935

   stack problem pare with 1918
   1918: infix expression to postfix expression
   1935: calculate postfix expression
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
	expr := []byte(getString(r))
	var maps []int = make([]int, N, N)

	for i := 0; i < N; i++ {
		maps[i] = getInt(r)
	}

	st := MakeStack() // for operand

	for i := 0; i < len(expr); i++ {
		switch {
		case expr[i] >= 'A' && expr[i] <= 'Z':
			st.Push(float64(maps[conv2Key(expr[i])]))

		case expr[i] == '+':
			b := st.Pop()
			a := st.Pop()
			st.Push(a + b)

		case expr[i] == '-':
			b := st.Pop()
			a := st.Pop()
			st.Push(a - b)

		case expr[i] == '*':
			b := st.Pop()
			a := st.Pop()
			st.Push(a * b)

		case expr[i] == '/':
			b := st.Pop()
			a := st.Pop()
			st.Push(a / b)

		default:
			panic("Unknown Value: " + string(expr[i]))
		}
	}

	fmt.Printf("%.2f\n", st.Pop())
	w.Flush()
}

func conv2Key(key byte) int {
	return int(key - 'A')
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

func (st *Stack) Push(value float64) {
	var node *Node
	if st.size == 0 {
		node = MakeNode(value, nil)
	} else {
		node = MakeNode(value, st.head)
	}
	st.head = node
	st.size++
}

func (st *Stack) Pop() float64 {
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

func (st *Stack) Top() float64 {
	if st.size == 0 {
		panic("There are no element in stack")
	} else {
		return st.head.value
	}
}

type Node struct {
	value float64
	next  *Node
}

func MakeNode(value float64, next *Node) *Node {
	return &Node{value, next}
}
