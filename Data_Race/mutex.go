package main

import (
    . "fmt"
    "runtime"
    "time"
    "sync"
   // "sync/atomic"
)

var i = 0

//var state = make(map[int]int)

var mutex = &sync.Mutex{}

func incrementing() {
	for j := 0; j < 999990; j++{
		mutex.Lock()
		i = i + 1
		mutex.Unlock()
	}
}

func decrementing() {
	for k := 0; k < 999999; k++{
		mutex.Lock()
		i = i - 1
		mutex.Unlock()
	}
}

func main() {
    Println("Value of i at start:", i)
    runtime.GOMAXPROCS(runtime.NumCPU())
	go incrementing()
	go decrementing()

    time.Sleep(1000*time.Millisecond)
    Println("The magic number is:", i)
}
