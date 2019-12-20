package main

import (
	"os"
	"bufio"
	"strconv"
	"strings"
	"fmt"
)

func main(){
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	
	var cnt int = 1

	for {
		strTkn := getStrTkn(r)

		L,_ := strconv.Atoi(strTkn[0])
		P,_ := strconv.Atoi(strTkn[1])
		V,_ := strconv.Atoi(strTkn[2])

		if L==0 && P==0 && V==0 {
			break
		}
		
		res := V/P*L
		if V%P > L {
			res += L
		} else {
			res += V%P
		}

		output := fmt.Sprintf("Case %d: %d\n", cnt, res)
		w.WriteString(output)
		cnt++
	}

	w.Flush()
}


func getStrTkn(r *bufio.Reader) []string{
	str,_ := r.ReadString('\n')
	str = strings.TrimSpace(str)
	strTkn := strings.Split(str," ")

	return strTkn
}