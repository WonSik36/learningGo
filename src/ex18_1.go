package main
 
// goroutine
// concurrent -> n goroutine, 1 cpu
// parallelism -> n goroutine, n cpu

import (
    "fmt"
    "time"
)
 
func say(s string) {
    for i := 0; i < 10; i++ {
        fmt.Println(s, "***", i)
    }
}
 
func main() {
    // Execute function synchronously
    say("Sync")
 
    // Execute function asynchronously
    go say("Async1")
    go say("Async2")
    go say("Async3")
 
    // wait 3 seconds
    time.Sleep(time.Second * 3)
}