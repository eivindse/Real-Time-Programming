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
	 
 

	// loop with switch
	for {
		 select {
		case num := <- add_number: //adding the next number in the channel que 
			i += num
		case cmd := <- control:
			switch cmd {
			case GetNumber:
				// update number
				number <- i
			case Exit:
				// end goroutine
				return
			}
		}
	 }
}

func incrementing(add_number chan<-int, finished chan<- int) {
	for j := 0; j<999995; j++ {
		add_number <- 1 // adding a +1 to the que
	}  
	//signal that the goroutine is finished
	finished <- 1
}

func decrementing(add_number chan<- int, finished chan<- int) {
	for y := 0; y<1000000; y++ {
		add_number <- -1 // adding a -1 to the que
	}
	//signal that the goroutine is finished
	finished <- 1
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	Println("Value of i at start:", i)
 
	 // Construct the required channels
	add_number := make(chan int)
	number := make(chan int)
	control := make(chan int)
	finished := make(chan int)
 
	// Spawn the required goroutines
	go incrementing(add_number, finished)
	go decrementing(add_number, finished)
	go number_server(add_number, control, number)

	// block on finished from both "worker" goroutines

	<- finished
	<- finished

	control<-GetNumber
	Println("The magic number is:", <- number)
	control<-Exit
}