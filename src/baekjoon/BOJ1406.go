/*
   baekjoon online judge
   problem number 1406
   https://www.acmicpc.net/problem/1406

   Make Deque Data Structure
*/

package main

import (
	"bufio"
	_ "fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	str := getString(r)
	dq := newDeque(str)

	N, _ := strconv.Atoi(getString(r))
	var strTkn []string
	for i := 0; i < N; i++ {
		strTkn = getStrTkn(r)

		switch {
		case strTkn[0] == "L":
			dq.moveCurPrev()
		case strTkn[0] == "D":
			dq.moveCurNext()
		case strTkn[0] == "B":
			dq.deleteCur()
		case strTkn[0] == "P":
			bytes := []byte(strTkn[1])
			dq.addNode(bytes[0])
		}
		// fmt.Printf("%c\n",dq.cur.value)
		// fmt.Print(dq.toString())
	}

	w.WriteString(dq.toString())
	w.Flush()
}

type node struct {
	value byte
	next  *node
	prev  *node
}

func newNode(value byte) *node {
	n := node{value, nil, nil}
	return &n
}

type deque struct {
	size int
	head *node
	tail *node
	cur  *node
}

func newDeque(str string) *deque {
	// head and tail is dummy node
	dq := deque{0, newNode(0), newNode(0), nil}
	dq.head.prev = dq.head
	dq.tail.next = dq.tail

	bytes := []byte(str)
	for i := 0; i < len(bytes); i++ {
		(&dq).addNode(bytes[i])
		// fmt.Printf("%c\n",dq.cur.value)
		// fmt.Print(dq.toString())
	}

	return &dq
}

func (dq *deque) addNode(value byte) {
	n := newNode(value)

	if dq.cur == dq.tail {
		panic("cur has same position with tail")
	}

	if dq.size == 0 {
		// step1
		n.prev = dq.head
		n.next = dq.tail
		// step2
		dq.head.next = n
		dq.tail.prev = n
	} else {
		// step1
		n.prev = dq.cur
		n.next = dq.cur.next
		// step2
		dq.cur.next.prev = n
		dq.cur.next = n
	}

	// step3
	dq.cur = n
	dq.size++
}

func (dq *deque) moveCurPrev() {
	if dq.cur == dq.tail {
		panic("cur has same position with tail")
	}

	dq.cur = dq.cur.prev
}

func (dq *deque) moveCurNext() {
	if dq.cur == dq.tail {
		panic("cur has same position with tail")
	}

	if dq.cur.next != dq.tail {
		dq.cur = dq.cur.next
	}
}

func (dq *deque) deleteCur() {
	if dq.cur == dq.tail {
		panic("cur has same position with tail")
	}

	if dq.size == 0 || dq.cur == dq.head {
		return
	}

	dq.cur.prev.next = dq.cur.next
	dq.cur.next.prev = dq.cur.prev
	dq.cur = dq.cur.prev
	dq.size--
}

func (dq *deque) toString() string {
	buf := []byte("")
	tmp := dq.head.next

	for i := 0; i < dq.size; i++ {
		buf = append(buf, tmp.value)
		tmp = tmp.next
	}

	buf = append(buf, '\n')
	return string(buf)
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
