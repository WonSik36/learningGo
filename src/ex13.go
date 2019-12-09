package main

import "fmt"
/* struct */
/* 
	In Go, a struct has only field data and no methods
*/

// Person means public person
type person struct{
	name string
	age int
}

type dict struct {
    data map[int]string
}

func main(){
    p1 := person{}
    p1.name = "Lee"
    p1.age = 10
	 
	var p2 person 
	p2 = person{"Bob", 20}
	p3 := person{name: "Sean", age: 50}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p3)

	dic := newDict() 	// call constructor
	dic.data[1] = "A"	// do not use -> like C, just use .
	fmt.Println(dic.data[1])
}

func newDict() *dict {
    d := dict{}
    d.data = map[int]string{}
    return &d //return pointer
}