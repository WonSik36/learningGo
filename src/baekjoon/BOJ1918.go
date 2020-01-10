/*
   baekjoon online judge
   problem number 1918
   https://www.acmicpc.net/problem/1918
   https://ghkvud2.tistory.com/67

   I used two stacks for operator, and operand
   But operand doesn't need stack
*/

package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	st := MakeStack()       // for operator:  +, -, *, /, (
	var res strings.Builder // for operand

	infix := []byte(getString(r))

	for i := 0; i < len(infix); i++ {
		switch infix[i] {
		case '(':
			st.Push(infix[i])

		case ')':
			for !st.Empty() {
				top := st.Pop()
				if top != '(' {
					res.WriteByte(top)
				} else {
					break
				}
			}

		case '+':
			fallthrough
		case '-':
			fallthrough
		case '*':
			fallthrough
		case '/':
			for !st.Empty() {
				if getPriority(st.Top()) >= getPriority(infix[i]) {
					res.WriteByte(st.Pop())
				} else {
					break
				}
			}
			st.Push(infix[i])

		default:
			res.WriteByte(infix[i])
		}
	}

	for !st.Empty() {
		res.WriteByte(st.Pop())
	}

	res.WriteByte('\n')

	w.WriteString(res.String())
	w.Flush()
}

func getPriority(c byte) int {
	switch c {
	case '(':
		return 0
	case '+':
		fallthrough
	case '-':
		return 1
	case '*':
		fallthrough
	case '/':
		return 2
	default:
		panic("Unknown Operator " + string(c))
	}
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

type Stack struct {
	size int
	head *Node
}

func MakeStack() *Stack {
	return &Stack{0, nil}
}

func (st *Stack) Push(value byte) {
	var node *Node
	if st.size == 0 {
		node = MakeNode(value, nil)
	} else {
		node = MakeNode(value, st.head)
	}
	st.head = node
	st.size++
}

func (st *Stack) Pop() byte {
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

func (st *Stack) Top() byte {
	if st.size == 0 {
		panic("There are no element in stack")
	} else {
		return st.head.value
	}
}

type Node struct {
	value byte
	next  *Node
}

func MakeNode(value byte, next *Node) *Node {
	return &Node{value, next}
}
