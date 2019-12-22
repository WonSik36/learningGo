/*
    baekjoon online judge
    problem number 
    https://www.acmicpc.net/problem/
*/

package main

import (
	"os"
	"bufio"
	"strconv"
	"strings"
	"sort"
)

func main(){
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	w.Flush()
}

func getString(r *bufio.Reader) string{
	str,_ := r.ReadString('\n')
	str = strings.TrimSpace(str)

	return str
}

func getStrTkn(r *bufio.Reader) []string{
	str,_ := r.ReadString('\n')
	str = strings.TrimSpace(str)
	strTkn := strings.Split(str," ")

	return strTkn
}

func printArr(arr []int, w *bufio.Writer) {
	for i:=0;i<len(arr);i++ {
		w.WriteString(strconv.Itoa(arr[i])+"\n")
	}
}