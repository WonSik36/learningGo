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
	
	var arr [9]int
	var sum int = 0

	for i:=0;i<9;i++ {
		str := getStr(r)
		arr[i],_ = strconv.Atoi(str)
		sum += arr[i]
	}

	var target int = sum - 100
	var i,j int

LOOP:
	for i=0;i<8;i++ {
		for j=i+1;j<9;j++ {
			if arr[i]+arr[j] == target {
				break LOOP
			}
		}	
	}

	var res []int
	var cnt int = 0

	for k:=0;k<9;k++ {
		if k==i || k==j {
			continue
		}

		res = append(res,arr[k])
		cnt++
	}

	sort.Ints(res)
	printArr(res,w)

	w.Flush()
}


func getStr(r *bufio.Reader) string{
	str,_ := r.ReadString('\n')
	str = strings.TrimSpace(str)

	return str
}

func printArr(arr []int, w *bufio.Writer) {
	for i:=0;i<len(arr);i++ {
		w.WriteString(strconv.Itoa(arr[i])+"\n")
	}
}