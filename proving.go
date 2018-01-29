package main

import (
    //. "fmt"
    . "runtime"
     "time"
)

/*
func inc(c chan int){
    for t := 0; t<1000; t++ {
		add_number <- 1
	} 
}
*/
func main(){
    runtime.GOMAXPROCS(runtime.NumCPU())
    
    number:= make(chan int)

    c := make(chan int)

    go func(){
        c <- 5
    }()
    
    go inc(c)
    
    time.Sleep(100*time.Millisecond)

}