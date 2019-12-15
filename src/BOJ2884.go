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

	hour,_ := strconv.Atoi(strToken[0])
	min,_ := strconv.Atoi(strToken[1])

	switch {
	case min>=45:
		min -= 45
	default:
		if hour == 0{
			hour = 23
		}else{
			hour--
		}

		min = 15 + min
	}

	n1 := strconv.Itoa(hour)
	n2 := strconv.Itoa(min)

	w.WriteString(n1+" "+n2)
	w.Flush()
}