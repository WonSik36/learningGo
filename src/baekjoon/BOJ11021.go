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

	// str,_ := r.ReadString('\n')
	// str = strings.TrimSpace(str)
	str := getString(r)

	N,_ := strconv.Atoi(str)

	for i:=1;i<=N;i++ {
		// str,_ = r.ReadString('\n')
		// str = strings.TrimSpace(str)
		// strTkn := strings.Split(str," ")

		strTkn := getStrTkn(r)

		n1,_ := strconv.Atoi(strTkn[0])
		n2,_ := strconv.Atoi(strTkn[1])

		res := n1+n2

		w.WriteString("Case #"+strconv.Itoa(i)+": "+strconv.Itoa(res)+"\n")
	}

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