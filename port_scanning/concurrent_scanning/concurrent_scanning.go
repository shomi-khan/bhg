package concurrent_scanning

import (
	"fmt"
	"net"
	"sync"
)

func checkOpenPort(port int) bool {
	address := fmt.Sprintf("localhost:%d", port)
	conn, err := net.Dial("tcp", address)
	if err == nil {
		conn.Close()
		return true
	}
	return false
}

func worker(pools chan int, wg *sync.WaitGroup) {
	for i := range pools {
		isOpen := checkOpenPort(i)
		if isOpen {
			fmt.Printf("Port %d is open\n", i)
		}
		wg.Done()
	}
}

func ScanPorts() {
	fmt.Println("Starting concurrent port scanning...")
	var wg sync.WaitGroup
	defer wg.Wait()

	pools := make(chan int, 100)
	for i := 0; i < cap(pools); i++ {
		go worker(pools, &wg)
	}

	for i := 1; i <= 65535; i++ {
		wg.Add(1)
		pools <- i
	}
}
