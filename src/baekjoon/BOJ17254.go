/*
	baekjoon online judge
	problem number 17254
	https://www.acmicpc.net/problem/17254
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

	arr := getIntArr(r)
	var nodeList []Node

	for i := 0; i < arr[1]; i++ {
		strTkn := getStrTkn(r)
		user, _ := strconv.Atoi(strTkn[0])
		time, _ := strconv.Atoi(strTkn[1])
		char := []byte(strTkn[2])[0]

		nodeList = append(nodeList, Node{user, time, char})
	}

	sort.Slice(nodeList, func(i, j int) bool {
		if nodeList[i].time < nodeList[j].time {
			return true
		} else if nodeList[i].time > nodeList[j].time {
			return false
		} else {
			if nodeList[i].user < nodeList[j].user {
				return true
			} else if nodeList[i].user > nodeList[j].user {
				return false
			}
		}
		return false
	})

	// printArr(nodeList)

	for i := 0; i < len(nodeList); i++ {
		w.WriteByte(nodeList[i].char)
	}

	w.WriteByte('\n')
	w.Flush()
}

type Node struct {
	user int
	time int
	char byte
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

func printArr(arr []Node) {
	for i := 0; i < len(arr); i++ {
		fmt.Printf("%d %d %c\n", arr[i].user, arr[i].time, arr[i].char)
	}
	fmt.Print("\n")
}
