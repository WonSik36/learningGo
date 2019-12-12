package main
 
// channel

/* 
	In an unbuffered channel, the sender is tied to the channel sending the data until one receiver receives the data.
	Using a Buffered Channel, even if the receiver is not ready to receive it, 
	you can send data as much as the specified buffer and continue doing other things.
*/
import "fmt"
 
func main() {
/*
	c := make(chan int)
  	c <- 1   // Deadlock because there is no incoming routine
  	fmt.Println(<-c) // Deadlock even if comments (because there is no separate Go routine)
*/
    ch := make(chan int, 1)
 
    // can send even if there is no recipient.
    ch <- 101
 
    fmt.Println(<-ch)
}