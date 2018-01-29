package main

import (
    "fmt"
    "runtime"
    "time"
)

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())

	ch := make(chan int)
    go ting(ch)

    time.Sleep(1000*time.Millisecond)
    fmt.Println(<-ch)
    fmt.Println(<-ch)
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}

func ting(ch chan int) {
    ch <- 1
    ch <- 5
}