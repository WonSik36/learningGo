package main

import (
	"os"
	"bufio"
	"strconv"
	"strings"
)

/*
	(a^b)%c by java

	public static long calcPow(int a, int b, int c){
        long result = 1;
        long multiply = a%c;

        while(b > 0){
            if(b%2 == 1){
                result *= multiply;
                result %= c;
            }
            multiply = multiply * multiply % c;
            b /= 2;
        }
        return result;
    }
*/

func main(){
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	
	strTkn := getStrTkn(r)

	a,_ := strconv.Atoi(strTkn[0])
	b,_ := strconv.Atoi(strTkn[1])
	c,_ := strconv.Atoi(strTkn[2])

	res := int(calcPow(a,b,c))
	w.WriteString(strconv.Itoa(res)+"\n")
	w.Flush()
}

func getStrTkn(r *bufio.Reader) []string{
	str,_ := r.ReadString('\n')
	str = strings.TrimSpace(str)
	strTkn := strings.Split(str," ")

	return strTkn
}

func calcPow(A int, B int, C int) int64 {
	a := int64(A)
	b := int64(B)
	c := int64(C)
	var res int64 = 1
	var mult int64 = a%c

	for b>0 {
		if b%2 == 1 {
			res *= mult
			res %= c
		}
		mult = mult * mult % c
		b /= 2
	}

	return res
}