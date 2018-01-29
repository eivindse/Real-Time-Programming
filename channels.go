// Use `go run foo.go` to run your program
 
package main
 
import (
	. "fmt"
	"runtime"
)

// Control signals
const (
	GetNumber = iota
	Exit
)

var number = 1
var i = 0
 func number_server(add_number <-chan int, control <-chan int, number chan<- int) {
	 
 

	// This for-select pattern is one you will become familiar with if you're using go "correctly".
	for {
		 select {
		case num := <- add_number:
			i += num
		case cmd := <- control:
			switch cmd {
			case GetNumber:
				// TODO: Send updated number
				number <- i
			case Exit:
				// TODO: End this goroutine gracefully
				return
			}
			// TODO: receive different messages and handle them correctly
			// You will at least need to update the number and handle control signals.
		 }
	 }
}

func incrementing(add_number chan<-int, finished chan<- int) {
	for j := 0; j<999995; j++ {
		add_number <- 1
	}  
	//TODO: signal that the goroutine is finished
	finished <- 1
}

func decrementing(add_number chan<- int, finished chan<- int) {
	for y := 0; y<1000000; y++ {
		add_number <- -1
	}
	//TODO: signal that the goroutine is finished
	finished <- 1
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	Println("Value of i at start:", i)
 
	 // TODO: Construct the required channels
	 // Think about wether the receptions of the number should be unbuffered, or buffered with a fixed queue size.
	add_number := make(chan int)
	number := make(chan int)
	control := make(chan int)
	finished := make(chan int)
 
	// TODO: Spawn the required goroutines
	go incrementing(add_number, finished)
	go decrementing(add_number, finished)
	go number_server(add_number, control, number)

	// TODO: block on finished from both "worker" goroutines

	<- finished
	<- finished

	control<-GetNumber
	Println("The magic number is:", <- number)
	control<-Exit
}