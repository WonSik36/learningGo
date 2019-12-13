package main

import(
	"os"
	"bufio"
)

func main(){
	str := 
`\    /\
 )  ( ')
(  /  )
 \(__)|`

	w := bufio.NewWriter(os.Stdout)
	w.Write([]byte(str))
	w.Flush()
}