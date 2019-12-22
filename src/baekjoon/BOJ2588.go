package main

import(
	"os"
	"bufio"
	"strconv"
	s "strings"
)

func main(){
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	
	str1,_ := r.ReadString('\n')
	str2,_ := r.ReadString('\n')

	str1 = s.TrimRight(str1,"\n")
	str2 = s.TrimRight(str2,"\n")
	str1 = s.TrimRight(str1,"\r")
	str2 = s.TrimRight(str2,"\r")

	n1,e := strconv.Atoi(str1)
	if e != nil{
		panic(e)
	}
	n2,e := strconv.Atoi(str2)
	if e != nil{
		panic(e)
	}

	res1 := (n2 % 10) * n1
	res2 := (n2 / 10 % 10) * n1
	res3 := (n2 / 100 % 10) * n1
	res4 := n1 * n2

	w.WriteString(strconv.Itoa(res1))
	w.WriteString("\n")
	w.WriteString(strconv.Itoa(res2))
	w.WriteString("\n")
	w.WriteString(strconv.Itoa(res3))
	w.WriteString("\n")
	w.WriteString(strconv.Itoa(res4))
	w.WriteString("\n")
	w.Flush()
}