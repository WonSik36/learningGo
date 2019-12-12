package main
 
// channel

import "fmt"
 
func main() {
    // make int type channel
    ch := make(chan int)

    go func() {
        ch <- 123   // send 123 to channel
    }()

    var i int
    i = <- ch  // get 123 from channel
    println(i)

    done := make(chan bool)
    go func() {
        for i := 0; i < 10; i++ {
            fmt.Println(i)
        }
        done <- true
    }()
 
    // Wait until the end of Goroutine above
    <-done
}