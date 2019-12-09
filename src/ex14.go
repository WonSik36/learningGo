package main

/* method */

type Rect struct {
    width, height int
}

/* call by value */
func (r Rect) area1() int {
	r.width++
    return r.width * r.height   
}
 
/* call by reference */
func (r *Rect) area2() int {
    r.width++
    return r.width * r.height
}

func main() {
    rect1 := Rect{10, 20}
    area1 := rect1.area1()
	println(area1, rect1.width)

	rect2 := Rect{10, 20}
    area2 := rect2.area2()
	println(area2, rect2.width)
}