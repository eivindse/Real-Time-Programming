package main
 
import (
    "fmt"
    "net"
    "os"
)

func main(){

}

func amIMaster(){//run if change in number of elevator
	//Listen to udp for living elevators ip
	//if one ip is lower than mine do nothing
	//if my is lowest I am master
	//return true if master
}

func livingElevator(){
	// listen to broadcast and count
	// count number of different ip
	return alive
}

/*struct elevator_Status{
	string ip
	int floor
	var a [5]int //direction, door, light
}*/

func master_Order(elevator_status[], order){
	//do algorithm
}

func que(){

}

func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}

func TCP(){

}

func UDPListen(){
	ServerConn, err := net.ListenPacket("udp", ":20013")
    CheckError(err)
    defer ServerConn.Close()
 
    buf := make([]byte, 1024)

    ipArray := make([])

    for {
        n,addr,err := ServerConn.ReadFrom(buf)
        fmt.Println("Received ",string(buf[0:n]), " from ",addr)

 
        if err != nil {
            fmt.Println("Error: ",err)
        }
        return buf, addr
    }

}

func UDPBroadcast(){
    Conn, err := net.Dial("udp", ":20013")
    CheckError(err)
    floot := (rand.Intn(100))

 
    defer Conn.Close()
    for {
        msg := fmt.Sprintf("Floor %d, going up", floot)
        buf := []byte(msg)
        _,err := Conn.Write(buf)
        if err != nil {
           fmt.Println(msg, err)
        }
        time.Sleep(time.Second*5)
    }
}