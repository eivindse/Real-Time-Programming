package main
 
import (
    "fmt"
    "net"
    "os"
)
 
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}
 
func main() {
    ServerConn, err := net.ListenPacket("udp", ":20013")
    CheckError(err)
    defer ServerConn.Close()
 
    buf := make([]byte, 1024)
 
    for {
        n,addr,err := ServerConn.ReadFrom(buf)
        fmt.Println(addr, ": ", string(buf[0:n]))
 
        if err != nil {
            fmt.Println("Error: ",err)
        } 
    }
}
