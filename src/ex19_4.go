package main
 
func main() {
    ch := make(chan int, 2)
 
    // transmit to channel
    ch <- 1
    ch <- 2
	
    // close channel
    close(ch)
 
    // first way
    // Continue to receive until it detects that the channel is closed
    /*
    for {
        if i, success := <-ch; success {
            println(i)
        } else {
            break
        }
    }
    */
 
    // second way
    // Same channel range statement as above
    for i := range ch {
        println(i)
    }
}