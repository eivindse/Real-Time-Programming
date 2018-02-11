package main

import (
    "fmt"
    "net"
    "os"
)

func handleRequest(conn net.Conn,thingsToDo chan<-int) {

    buf := make([]byte, 1024)
  
    _, err := conn.Read(buf) 
    
    if err != nil {
      fmt.Println("Error reading:", err.Error())
    }else{
      fmt.Println("Recieving: " + string(buf[:]))
    }
  
    conn.Write([]byte(buf))
    conn.Write([]byte("Thank you"+"\n"))
    thingsToDo <- 2
    conn.Close()
  }

  func sendStuff(conn net.Conn, msg int) {
    //thingsToDo = false
    fmt.Println(msg)

    buf := make([]byte, 1024)
  
    _, err := conn.Read(buf) 
    
    if err != nil {
      fmt.Println("Error reading:", err.Error())
    }else{
      //fmt.Println("Recieving: " + string(buf[:]))
    }
  
    conn.Write([]byte(buf))
    conn.Write([]byte("Thank you"+"\n"))
    
    conn.Close()
  }


func main() {
    thingsToDo := make(chan int)

    l, err := net.Listen("tcp", "localhost"+":"+"3333")
    if err != nil {
        fmt.Println("Error listening:", err.Error())
        os.Exit(1)
    }
    defer l.Close()

    for {
        conn, err := l.Accept()
        if err != nil {
            fmt.Println("Error accepting: ", err.Error())
            thingsToDo<-1
            os.Exit(1)
        }
        go handleRequest(conn,thingsToDo)

        select{
        case msg := <-thingsToDo:
                go sendStuff(conn, msg)
        }
    }

}
/*

func main() {
	//TCP-setup
	raddr, err := net.ResolveTCPAddr("tcp", "129.241.187.136:33546")
	if err != nil {log.Fatal(err)}

	socket, err := net.DialTCP("tcp", nil, raddr)
	if err != nil {log.Fatal(err)}
	
	//declare variables
	var msg = "Connect to: 129.241.187.142:20014\x00"
	
	//make listener
	laddr, err := net.ResolveTCPAddr("tcp", "129.241.187.142:20014")
	if err != nil {log.Fatal(err)}
	
	listener, err := net.ListenTCP("tcp", laddr)
	if err != nil {log.Fatal(err)}
	
	//server connect back
	_, err = socket.Write([]byte(msg))
	if err != nil {log.Fatal(err)}
	
	socket_connect, err := listener.AcceptTCP()
	if err != nil {log.Fatal(err)}
	
	doneChannel := make(chan bool, 1);
	
	go ReadFromWriteTo(socket, doneChannel)
	go accept_connection(socket_connect, doneChannel)
	
	<- doneChannel
}*/


