package main

import (
    "fmt"
    "net"
    "os"
    "time"
)

func isPortOpen(ip string, port int) bool {
    address := fmt.Sprintf("%s:%d", ip, port)
    conn, err := net.DialTimeout("tcp", address, 2*time.Second)
    if err != nil {
        return false
    }
    defer conn.Close()
    return true
}

func main() {
    if len(os.Args) != 3 {
        fmt.Println("Usage: go run check_port.go <IP_ADDRESS> <PORT>")
        return
    }

    ip := os.Args[1]
    var port int
    _, err := fmt.Sscanf(os.Args[2], "%d", &port)
    if err != nil || port <= 0 || port > 65535 {
        fmt.Println("Invalid port number. Please provide a valid port between 1 and 65535.")
        return
    }

    if isPortOpen(ip, port) {
        fmt.Printf("Port %d is open on %s\n", port, ip)
    } else {
        fmt.Printf("Port %d is closed on %s\n", port, ip)
    }
}
