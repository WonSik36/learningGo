package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"strconv"
)

func main(){
	var pNum int

	// get problem number
	fmt.Print("Problem number: ")
	_,err := fmt.Scan(&pNum)
	checkError(err)

	// get directory path of input problem number & template
	dir,err := os.Getwd() 
	checkError(err)
	temp := fmt.Sprintf(dir+"\\src\\baekjoon\\template.go") 
	dir = fmt.Sprintf(dir+"\\src\\baekjoon\\BOJ"+strconv.Itoa(pNum)+".go") 
	fmt.Println(dir)

	// check if problem number is dupliacted
	if _, err := os.Stat(dir); !os.IsNotExist(err) {
		fmt.Println("File already exist")
		return
	} 

	// get template file
	dat,err := ioutil.ReadFile(temp)
	checkError(err)

	// get destination file
	err = ioutil.WriteFile(dir,dat,os.FileMode(0644))
	checkError(err)

	fmt.Println("Make file succeed")
}

func checkError(err error){
	if err != nil {
		panic(err)
	}
}