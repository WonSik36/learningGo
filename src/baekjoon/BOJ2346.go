/*
    baekjoon online judge
    problem number 2346
	https://www.acmicpc.net/problem/2346

	Make Deque Data Structure
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

	N, _ := strconv.Atoi(getString(r))
	dq := newDeque()

	// insert input to deque
	strTkn := getStrTkn(r)
	for i := 0; i < N; i++ {
		num, _ := strconv.Atoi(strTkn[i])
		dq.addNode2Tail(i+1, num) // key, value
	}

	// dq.print()

	// remove node and print its value
	dq.setPos2Head()
	// fmt.Printf("pos: (%d,%d)\n",dq.pos.key,dq.pos.value)
	k := 0
	for dq.size > 0 {
		k, _ = dq.removeAndMovePos()
		// fmt.Printf("deque size: %d\n",dq.size)
		w.WriteString(strconv.Itoa(k) + " ")
	}

	w.Flush()
}

type node struct {
	key   int
	value int
	next  *node
	prev  *node
}

type deque struct {
	size int
	head *node
	tail *node
	pos  *node
}

func newNode(key int, value int) *node {
	n := node{key, value, nil, nil}
	return &n
}

func (n *node) nextNode(step int) *node {
	var tmp *node = n
	if step == 0 {
	} else if step > 0 {
		for i := 0; i < step; i++ {
			tmp = tmp.next
			if tmp == n {
				i--
			}
		}
	} else {
		step = step * -1
		for i := 0; i < step; i++ {
			tmp = tmp.prev
			if tmp == n {
				i--
			}
		}
	}

	return tmp
}

func newDeque() *deque {
	dq := deque{0, nil, nil, nil}
	return &dq
}

func (dq *deque) addNode2Tail(key int, value int) {
	// n is pointer of node
	n := newNode(key, value)

	// size is 0
	if dq.size == 0 {
		dq.head = n
		dq.tail = n
		n.next = n
		n.prev = n
		// size is larger than 0
	} else {
		n.next = dq.head
		n.prev = dq.tail
		dq.head.prev = n
		dq.tail.next = n

		dq.tail = n
	}

	dq.size++
}

func (dq *deque) setPos2Head() {
	dq.pos = dq.head
}

func (dq *deque) removeAndMovePos() (int, int) {
	if dq.pos == nil {
		panic("pos is not initialized")
	}

	if dq.size <= 0 {
		panic("deque size is less than or equal to 0")
	} else if dq.size == 1 {
		k, v := dq.pos.key, dq.pos.value

		dq.size = 0
		dq.head = nil
		dq.tail = nil
		dq.pos = nil

		return k, v
	}

	// get key, value and next pos
	k := dq.pos.key
	v := dq.pos.value
	nextPos := dq.pos.nextNode(v)
	// fmt.Printf("next pos: (%d,%d)\n",nextPos.key,nextPos.value)

	// update previous and next node of pos
	dq.pos.prev.next = dq.pos.next
	dq.pos.next.prev = dq.pos.prev

	// update size and pos
	dq.size--
	dq.pos = nextPos
	// fmt.Printf("pos: (%d,%d)\n",dq.pos.key,dq.pos.value)
	return k, v
}

func (dq *deque) print() {
	tmp := dq.head

	for tmp != dq.tail {
		fmt.Printf("%d:%d ", tmp.key, tmp.value)
		tmp = tmp.next
	}

	fmt.Printf("%d:%d\n", tmp.key, tmp.value)
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
