package main
 
// switch

import "fmt"
 
func main() {
    check(0)
}
 
func check(val int) {
    switch {
    case val<1:
        fmt.Println("1 이하")
        fallthrough
    case val<2:
        fmt.Println("2 이하")
        fallthrough
    case val<3:
        fmt.Println("3 이하")
        fallthrough
    default:
        fmt.Println("default 도달")
    }
}