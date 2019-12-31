/*
    baekjoon online judge
    problem number 3078
	https://www.acmicpc.net/problem/3078

	Time Out

	Line Sweeping(Deprecated)
	https://www.acmicpc.net/blog/view/25
	https://mygumi.tistory.com/294
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	strTkn := getStrTkn(r)
	N, _ := strconv.Atoi(strTkn[0])
	K, _ := strconv.Atoi(strTkn[1])
	var arr []Point = make([]Point, N, N)

	// get slice
	for i := 0; i < N; i++ {
		str := getString(r)
		arr[i] = Point{i, len(str)}
	}

	// sort slice by length
	sort.Slice(arr, func(i, j int) bool {
		return arr[i].length < arr[j].length
	})
	// printArr(arr)

	var res int64 = 0
	// loop1
	sIdx := 0
	for sIdx < N-1 {
		// loop2
		// get range of same length in slice
		eIdx := getMaxIdxByLength(arr[sIdx].length, arr)

		// copy the range of slice to temporary slice
		// element in copied slice has same length
		tmpArr := arr[sIdx : eIdx+1]

		// sort slice by rank
		sort.Slice(tmpArr, func(i, j int) bool {
			return tmpArr[i].rank < tmpArr[j].rank
		})

		// printArr(tmpArr)

		// loop3
		// get best friend (which has same length and gap of rank is same or under than K)
		for i := 0; i < len(tmpArr)-1; i++ {
			for j := i + 1; j < len(tmpArr); j++ {
				if (tmpArr[j].rank - tmpArr[i].rank) <= K {
					res++
				}
			}
		}

		sIdx += eIdx + 1
	}

	w.WriteString(fmt.Sprintf("%d\n", res))
	w.Flush()
}

type Point struct {
	rank   int // ranking of score, x coordinate
	length int // length of name
}

func (p *Point) toString() string {
	return fmt.Sprintf("rank: %d, length: %d", p.rank, p.length)
}

func getMaxIdxByLength(target int, arr []Point) int {
	var left, right int = 0, len(arr) - 1
	var mid int
	var res int

	for left <= right {
		mid = (left + right) / 2

		if arr[mid].length > target {
			right = mid - 1
		} else if arr[mid].length < target {
			left = mid + 1
		} else {
			res = mid
			left = mid + 1
		}
	}

	return res
}

func getMaxIdxByRank(target int, arr []Point) int {
	var left, right int = 0, len(arr) - 1
	var mid int
	var res int

	for left <= right {
		mid = (left + right) / 2

		if arr[mid].rank > target {
			right = mid - 1
		} else if arr[mid].rank < target {
			left = mid + 1
		} else {
			res = mid
			left = mid + 1
		}
	}

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

func printArr(arr []Point) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i].toString())
	}
	fmt.Println()
}
