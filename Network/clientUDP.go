package main
 
import (
    "fmt"
    "net"
    "time"
    "math/rand"
)
 
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}
 
func main() {
    Conn, err := net.Dial("udp", ":44343")
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
        time.Sleep(time.Millisecond*400)
    }
}