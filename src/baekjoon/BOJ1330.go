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
	
	strToken := s.Split(str," ")

	n1,_ := strconv.Atoi(strToken[0])
	n2,_ := strconv.Atoi(strToken[1])

	switch {
	case n1 == n2:
		w.WriteString("==\n")
	case n1 > n2:
		w.WriteString(">\n")
	case n1 < n2:
		w.WriteString("<\n")
	}

	w.Flush()
}