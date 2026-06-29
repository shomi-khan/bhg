package concurrent_scanning

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func checkOpenPort(port int) bool {
	address := fmt.Sprintf("localhost:%d", port)

	conn, err := net.DialTimeout("tcp", address, 500*time.Millisecond)
	if err != nil {
		return false
	}
	defer conn.Close()

	return true
}

func worker(pools <-chan int, results chan<- int, wg *sync.WaitGroup) {
	for port := range pools {
		if checkOpenPort(port) {
			results <- port
		}
		wg.Done()
	}
}

func ScanPorts() {
	fmt.Println("Starting concurrent port scanning...")
	start := time.Now()

	const totalPorts = 65535
	const workerCount = 100

	var wg sync.WaitGroup

	pools := make(chan int, workerCount)
	results := make(chan int, workerCount)

	// Start workers
	for i := 0; i < workerCount; i++ {
		go worker(pools, results, &wg)
	}

	// Send ports to workers
	for port := 1; port <= totalPorts; port++ {
		wg.Add(1)
		pools <- port
	}

	// No more jobs
	close(pools)

	// Close results after all workers finish
	go func() {
		wg.Wait()
		close(results)
	}()

	var openPorts []int

	// Receive results
	for port := range results {
		openPorts = append(openPorts, port)
	}

	fmt.Println("Opened ports:", openPorts)
	fmt.Println("Time taken:", time.Since(start))
}
