/*
	baekjoon online judge
	problem number 11651
	https://www.acmicpc.net/problem/11651

	I firstly used pointer to swap but go lang doesn't support pass by reference of slice
	So I referenced below site
	https://www.golangprograms.com/golang-program-for-implementation-of-quick-sort.html
	Referenced site uses sub-slice to change value of slice

	There are four impressive points.
	1. function can be argument value
		type ComparePoint func(p1, p2 *Point) int
		func quickSort(arr []*Point, compare ComparePoint) []*Point{

	2. no need to use swap function
		arr[left], arr[right] = arr[right], arr[left]

	3. new version of quicksort

	4. use sub-slice
		quickSort(arr[:left], compare)
*/

package main

import (
	"bufio"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	N := getInt(r)
	arr := make([]*Point, N, N)
	for i := 0; i < N; i++ {
		tmp := getIntArr(r)
		arr[i] = newPoint(tmp[0], tmp[1])
	}

	quickSort(arr, func(p1, p2 *Point) int {
		if p1.y < p2.y {
			return -1
		} else if p1.y > p2.y {
			return 1
		} else {
			if p1.x < p2.x {
				return -1
			} else if p1.x > p2.x {
				return 1
			} else {
				return 0
			}
		}
	})

	printArr(arr, w)
	w.Flush()
}

func quickSort(arr []*Point, compare ComparePoint) {
	// arr length is 0 or 1
	if len(arr) < 2 {
		return
	}

	left, right := 0, len(arr)-1

	pivot := rand.Int() % len(arr)

	arr[pivot], arr[right] = arr[right], arr[pivot]

	for i, _ := range arr {
		if compare(arr[i], arr[right]) == -1 {
			arr[left], arr[i] = arr[i], arr[left]
			left++
		}
	}

	arr[left], arr[right] = arr[right], arr[left]

	quickSort(arr[:left], compare)
	quickSort(arr[left+1:], compare)
}

type Point struct {
	x int
	y int
}

func newPoint(x, y int) *Point {
	p := Point{x, y}
	return &p
}

type ComparePoint func(p1, p2 *Point) int

type PointSlice []*Point

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

func printArr(arr []*Point, w *bufio.Writer) {
	var sb strings.Builder

	for i := 0; i < len(arr); i++ {
		sb.WriteString(strconv.Itoa(arr[i].x))
		sb.WriteString(" ")
		sb.WriteString(strconv.Itoa(arr[i].y))
		sb.WriteString("\n")
	}

	w.WriteString(sb.String())
}
