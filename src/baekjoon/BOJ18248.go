/*
	baekjoon online judge
	problem number 18248
	https://www.acmicpc.net/problem/18248

	Sort
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

var N int
var M int

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	var personList []*Person

	arr := getIntArr(r)
	N = arr[0]
	M = arr[1]

	for i := 0; i < N; i++ {
		arr = getIntArr(r)
		p := newPerson(arr)

		personList = append(personList, p)
	}

	sort.Slice(personList, func(i, j int) bool {
		return personList[i].cnt < personList[j].cnt
	})

	// printArr(personList)

	var isValid bool = true
	for i := 0; i < M; i++ {
		var flag bool = false

		for j := 0; j < N; j++ {
			if !personList[j].check[i] && flag {
				isValid = false
				break
			}

			if personList[j].check[i] {
				flag = true
			}
		}

		if isValid {
			break
		}
	}

	if isValid {
		w.WriteString("YES\n")
	} else {
		w.WriteString("NO\n")
	}

	w.Flush()
}

func newPerson(arr []int) *Person {
	var cnt int = 0
	var check []bool = make([]bool, M, M)

	for i := 0; i < M; i++ {
		if arr[i] == 1 {
			check[i] = true
			cnt++
		}
	}

	p := Person{cnt, check}
	return &p
}

type Person struct {
	cnt   int
	check []bool
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

func printArr(arr []*Person) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("cnt: %d, ", arr[i].cnt)
		for j := 0; j < len(arr[i].check); j++ {
			if arr[i].check[j] {
				fmt.Print("1 ")
			} else {
				fmt.Print("0 ")
			}
		}
		fmt.Println()
	}
	fmt.Print("\n")
}
