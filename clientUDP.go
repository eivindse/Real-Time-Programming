package main
 
import (
    "fmt"
    "net"
    "time"
)
 
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
    }
}
 
func main() {
    Conn, err := net.Dial("udp", ":20013")
    CheckError(err)
 
    defer Conn.Close()
    for {
        msg := "hei you"
        buf := []byte(msg)
        _,err := Conn.Write(buf)
        if err != nil {
           fmt.Println(msg, err)
        }
        time.Sleep(time.Second)
    }
}
