package concurrent_scanning

import (
	"fmt"
	"net"
	"sync"
)

func ScanPorts() {
	fmt.Println("Starting concurrent port scanning...")
	var wg sync.WaitGroup
	defer wg.Wait()

	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			address := fmt.Sprintf("localhost:%d", port)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				return
			}
			conn.Close()
			fmt.Printf("Port %d is open\n", port)
		}(i)
	}
}
