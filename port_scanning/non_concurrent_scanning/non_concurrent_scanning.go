package non_concurrent_scanning

import (
	"fmt"
	"net"
	"time"
)

func ScanPorts() {
	start := time.Now()
	fmt.Println("Starting non-concurrent port scanning...")
	for port := 1; port <= 65535; port++ {
		address := fmt.Sprintf("localhost:%d", port)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			continue
		}
		conn.Close()
		fmt.Printf("Port %d is open\n", port)
	}
	processingTime := time.Since(start)
	fmt.Println("Total taken time, ", processingTime)
}
