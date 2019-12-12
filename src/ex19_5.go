package main
 
/*
	Go's select statement provides the ability to execute a prepared (sending data) channel while waiting for multiple channels.
*/

import "time"
 
func main() {
    done1 := make(chan bool)
    done2 := make(chan bool)
 
    go run1(done1)
    go run2(done2)
 
EXIT:
    for {
        select {
        case <-done1:
            println("complete run1")
 
        case <-done2:
            println("complete run2")
            break EXIT	// In Go, the break statement moves to that label and executes the statement following the loop it exited.
        }
    }
}
 
func run1(done chan bool) {
    time.Sleep(1 * time.Second)
    done <- true
}
 
func run2(done chan bool) {
    time.Sleep(2 * time.Second)
    done <- true
}