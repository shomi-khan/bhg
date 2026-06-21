package main

import (
	"fmt"
	"net"
)

func main() {
	for port := 1; port <= 65535; port++ {
		address := fmt.Sprintf("localhost:%d", port)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			continue
		}
		conn.Close()
		fmt.Printf("Port %d is open\n", port)
	}
}
