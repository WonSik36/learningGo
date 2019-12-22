/*
    baekjoon online judge
    problem number 2104
	https://www.acmicpc.net/problem/2104

	reference: 
	https://pasudo123.tistory.com/284
	https://yangorithm.tistory.com/51
	high reference level
*/

package main

import (
	"os"
	"bufio"
	"strconv"
	"strings"
)

func main(){
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	str := getString(r)
	num,_ := strconv.Atoi(str)

	var arr []int = make([]int, num)
	strTkn := getStrTkn(r)
	for i:=0;i<num;i++ {
		arr[i],_ = strconv.Atoi(strTkn[i])
	}

	var res int64 = getPartArr(arr,0,len(arr)-1)

	w.WriteString(strconv.FormatInt(res,10)+"\n")
	w.Flush()
}

func getPartArr(arr []int, left int, right int) int64 {
	if left == right {
		return int64(arr[left]*arr[left])
	}

	var mid int = (left+right)/2

	var sum int64 = int64(arr[mid])
	var min int = arr[mid]
	var max int64 = sum * int64(min)
	var idxL int = mid-1
	var idxR int = mid+1

	for left <= idxL || idxR <= right {
		if idxR <= right && (idxL < left || arr[idxL] < arr[idxR]){		// key parts
			sum += int64(arr[idxR])
			min = Min(min, arr[idxR])
			idxR++
		} else if left <= idxL && (right < idxR || arr[idxL] >= arr[idxR]) { // key parts
			sum += int64(arr[idxL])
			min = Min(min, arr[idxL])
			idxL--
		}else{
			panic("assertion error")
		}

		max = Max(max, sum*int64(min))		// key parts
	}

	max = Max(max, getPartArr(arr,left,mid))
	max = Max(max, getPartArr(arr,mid+1,right))
	
	return max
}

func Max(a int64, b int64) int64 {
	if a>b {
		return a
	} else {
		return b
	}
}

func Min(a int, b int) int {
	if a<b {
		return a
	} else {
		return b
	}
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