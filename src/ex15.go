package main

/* interface */

import "fmt"
import "math"

// Shape
type Shape interface {
    area() float64
    perimeter() float64
}

// Rect
type Rect struct {
    width, height float64
}
 
// Circle
type Circle struct {
    radius float64
}
 
// Shape interface implementation for Rect types
func (r Rect) area() float64 { return r.width * r.height }
func (r Rect) perimeter() float64 {
     return 2 * (r.width + r.height)
}
 
// Shape Interface Implementation for Circle Type
func (c Circle) area() float64 { 
    return math.Pi * c.radius * c.radius
}
func (c Circle) perimeter() float64 { 
    return 2 * math.Pi * c.radius
}

func main() {
    r := Rect{10., 20.}
    c := Circle{10}
 
	showArea(r, c)
	
	/* empty interface */
	/* An empty interface can be thought of as an object in C # and Java, and a void * as in C / C ++. */
	var x interface{}
	x = 1 
	i := x
	j := x.(int)
	println(i)  // pointer address
    println(j)	// 1

    x = "Tom"
 
    printIt(x)	// Tom
}
 
func showArea(shapes ...Shape) {
    for _, s := range shapes {
        a := s.area()
        println(a)
    }
}

func printIt(v interface{}) {
    fmt.Println(v) //Tom
}