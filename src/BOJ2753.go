package main

import (
	"os"
	"bufio"
	"strconv"
	s "strings"
)

func main(){
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)

	str,_ := r.ReadString('\n')
	str = s.TrimSpace(str)
	num,_ := strconv.Atoi(str)

	switch {
	case num%400 == 0:
		w.WriteString("1\n")
	case num%100 == 0:
		w.WriteString("0\n")
	case num%4 == 0:
		w.WriteString("1\n")
	default:
		w.WriteString("0\n")
	}

	w.Flush()
}