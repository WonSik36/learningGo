package main
 
// goroutine
// concurrent -> n goroutine, 1 cpu
// parallelism -> n goroutine, n cpu

import (
    "fmt"
	"sync"
	"runtime"
)
 
func main() {
	// using 2 cpu -> parallelism
    runtime.GOMAXPROCS(2)

    // Create WaitGroup. Waiting for two Go routines.
    var wait sync.WaitGroup
    wait.Add(2)
 
    // Goroutine with anonymous functions
    go func() {
        defer wait.Done() // Call .Done () when finished
        fmt.Println("Hello")
    }()
 
    // Passing parameters to anonymous functions
    go func(msg string) {
		fmt.Println(msg)
        wait.Done()
    }("Hi")
 
	wait.Wait() // Wait until all Go routines are over
	println("go rountine end")
}