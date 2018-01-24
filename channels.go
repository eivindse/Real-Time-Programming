package main

import (
    . "fmt"
    "runtime"
    "time"
)

var i = 0
var x = 0



func incrementing(channel1 chan int, channel2 chan int) {
    
    for j := 0; j < 997; j++{
        select{
        case <- channel2:
            Println("i++")
             i = i + 1
             channel2<-1
        default:
            
        }     
    }
}

func decrementing(channel1 chan int, channel2 chan int) {
   
    for k := 0; k < 999; k++{
        select{
        case <- channel2:
            i = i - 1
            channel2<-1
        default:
            
        }
	}
}

func main() {
    Println("Value of i at start:", i)
    runtime.GOMAXPROCS(runtime.NumCPU())

    channel1 := make(chan int)
    channel2 := make(chan int)
    go func(){
        channel1 <- 1
        channel2 <- 1
        channel2 <- 1
        channel2 <- 1
    }()
    Println(<-channel2)
    go func(){ 
        channel2 <- 1
        channel2 <- 1
    }()

	go incrementing(channel1, channel2)
	go decrementing(channel1, channel2)

    time.Sleep(1000*time.Millisecond)
    Println("The magic number is:", i)
}
