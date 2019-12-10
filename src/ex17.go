package main
 
// defer, panic, recover

import (
	"os"
	"fmt"
)

func main() {
    openFile("Invalid.txt")
    println("Done") // it works because panic was recovered by recover()
}
 
func openFile(fn string) {
	f, err := os.Open(fn)
	defer f.Close()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("OPEN ERROR", r)
		}
	}()

    if err != nil {
		panic(err)
	}
	
	defer println("defer")	// it doesn't work because it was not registered
}