package main

import (
    . "fmt"
    "runtime"
    "time"
)

var i = 0

func incrementing() {
    for j := 0; j < 999999; j++{
		i = i + 1;
	}
}

func decrementing() {
    for k := 0; k < 999999; k++{
		i = i - 1;
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
