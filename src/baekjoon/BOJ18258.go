/*
	baekjoon online judge
	problem number 18258
	https://www.acmicpc.net/problem/18258

	make Queue
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
	q := makeQueue()
	for i := 0; i < N; i++ {
		// printQueue(q)
		tkn := getStrTkn(r)

		switch {
		case tkn[0] == "push":
			v, _ := strconv.Atoi(tkn[1])
			q.add(v)
		case tkn[0] == "pop":
			w.WriteString(strconv.Itoa(q.poll()))
			w.WriteString("\n")
		case tkn[0] == "size":
			w.WriteString(strconv.Itoa(q.size))
			w.WriteString("\n")
		case tkn[0] == "empty":
			if q.isEmpty() {
				w.WriteString("1\n")
			} else {
				w.WriteString("0\n")
			}
		case tkn[0] == "front":
			w.WriteString(strconv.Itoa(q.front()))
			w.WriteString("\n")
		case tkn[0] == "back":
			w.WriteString(strconv.Itoa(q.back()))
			w.WriteString("\n")
		}
	}

	w.Flush()
}

type Queue struct {
	head *Node
	tail *Node
	size int
}

func makeQueue() *Queue {
	head := makeNode(0)
	tail := makeNode(0)
	head.next = tail
	tail.prev = head
	return &Queue{head, tail, 0}
}

func (q *Queue) add(value int) {
	node := makeNode(value)

	node.prev = q.tail.prev
	node.next = q.tail
	q.tail.prev.next = node
	q.tail.prev = node

	q.size++
}

func (q *Queue) poll() int {
	if q.isEmpty() {
		return -1
	}

	node := q.head.next

	node.next.prev = q.head
	q.head.next = node.next
	q.size--

	return node.value
}

func (q *Queue) isEmpty() bool {
	if q.size == 0 {
		return true
	} else {
		return false
	}
}

func (q *Queue) front() int {
	if q.isEmpty() {
		return -1
	} else {
		return q.head.next.value
	}
}

func (q *Queue) back() int {
	if q.isEmpty() {
		return -1
	} else {
		return q.tail.prev.value
	}
}

func printQueue(q *Queue) {
	fmt.Printf("size: %d\n", q.size)

	node := q.head.next
	for node != q.tail {
		fmt.Printf("%d ", node.value)
		node = node.next
	}
	fmt.Println()
}

type Node struct {
	prev  *Node
	next  *Node
	value int
}

func makeNode(value int) *Node {
	return &Node{nil, nil, value}
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
