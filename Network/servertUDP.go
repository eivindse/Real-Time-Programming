package main
 
import (
    "fmt"
    "net"
    "os"
    "runtime"
    "time"
)
 
func CheckError(err error) {
    if err  != nil {
        fmt.Println("Error: " , err)
        os.Exit(0)
    }
}
func livingElevator(){
    // listen to broadcast and count


    type elevator struct {
        msg string
        ipaddr string
    }

    var ip_array []elevator

    for i := 0; i < 10; i++{
        buf, addr := UDPListen() 
        add := true
        for i := 0; i < len(ip_array); i++{
            tmp_ip := ip_array[i]
            address :=  addr.String()
            if(tmp_ip.ipaddr == address){
                add = false
            }
        }
        ip_new := elevator{string(buf), addr.String()}
          if add{
            ip_array = append(ip_array, ip_new)
          }
    }
    fmt.Println(len(ip_array))

	// count number of different ip
	//return alive
}

func UDPListen()([]byte, net.Addr){
    ServerConn, err := net.ListenPacket("udp", ":44343")
    CheckError(err)
    defer ServerConn.Close()
 
    buf := make([]byte, 1024)

    for {
        _,addr,err := ServerConn.ReadFrom(buf)
        //fmt.Println(addr, ": ", string(buf[0:n]))
 
        if err != nil {
            fmt.Println("Error: ",err)
        }        
        return buf, addr
    }
}

func main() {
    runtime.GOMAXPROCS(runtime.NumCPU())

    livingElevator()
    fmt.Println("Sleeping...")
    time.Sleep(time.Second*5)
    livingElevator()
    fmt.Println("Sleeping...")
    time.Sleep(time.Second*5)
    livingElevator()
    //buffer, adresse := UDPListen()
    //n := len(buffer)
    //fmt.Println(adresse, string(buffer[0:n]))

}
