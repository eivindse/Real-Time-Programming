package main

import (
    . "fmt"
    "runtime"
    "time"
)

func main(){

    runtime.GOMAXPROCS(runtime.NumCPU())

    c := make(chan int)

    c <- 1

    Println(<-c)

    time.Sleep(100*time.Millisecond)

}
