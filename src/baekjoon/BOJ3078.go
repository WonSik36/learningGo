/*
    baekjoon online judge
    problem number 3078
	https://www.acmicpc.net/problem/3078

	Using Queue for length
	https://blog.junu.dev/30
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

	strTkn := getStrTkn(r)
	N, _ := strconv.Atoi(strTkn[0])
	K, _ := strconv.Atoi(strTkn[1])
	var q []Queue = make([]Queue, 21, 21)
	var res int64 = 0

	// get student info
	for i := 0; i < N; i++ {
		// PrintArrayOfQueue(q)
		length := len(getString(r))

		for !q[length].Empty() {
			if (i - q[length].Peek().value) <= K {
				res += int64(q[length].size)
				break
			} else {
				q[length].Pop()
			}
		}

		q[length].Push(i)
	}

	fmt.Printf("%d\n", res)
}

type Node struct {
	value int
	next  *Node
}

func NewNode(value int) *Node {
	return &Node{value, nil}
}

type Queue struct {
	size int
	head *Node
	tail *Node
}

func NewQueue() Queue {
	return Queue{0, nil, nil}
}

func (q *Queue) Empty() bool {
	return q.size == 0
}

func (q *Queue) Push(value int) {
	n := NewNode(value)

	if q.size == 0 {
		q.head = n
		q.tail = n
	} else {
		q.tail.next = n
		q.tail = n
	}

	q.size++
}

func (q *Queue) Peek() *Node {
	if q.size == 0 {
		panic("Queue is empty")
	}

	return q.head
}

func (q *Queue) Pop() *Node {
	if q.size == 0 {
		panic("Queue is empty")
	}

	ret := q.head

	if q.size == 1 {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
	}

	q.size--
	return ret
}

func (q *Queue) Print() {
	for cur := q.head; cur != nil; cur = cur.next {
		fmt.Printf("%d ", cur.value)
	}
	fmt.Println()
}

func PrintArrayOfQueue(arrQ []Queue) {
	for i := 0; i < len(arrQ); i++ {
		fmt.Printf("Queue: %d\n", i)
		arrQ[i].Print()
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

func printArr(arr []int, w *bufio.Writer) {
	for i := 0; i < len(arr); i++ {
		w.WriteString(strconv.Itoa(arr[i]) + "\n")
	}
}
