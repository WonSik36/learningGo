package main

/*
	When a channel is also passed a function parameter, 
	it usually passes a channel that sends and receives both, 
	but it is also possible to specify whether to send or receive only on that channel.
*/
 
import "fmt"
 
func main() {
    ch := make(chan string, 1)
    sendChan(ch)
    receiveChan(ch)
}

// Transmission channel
func sendChan(ch chan<- string) {
	ch <- "Data"
    // x := <-ch // error
}
 
// Receiving channel
func receiveChan(ch <-chan string) {
    data := <-ch
    fmt.Println(data)
}