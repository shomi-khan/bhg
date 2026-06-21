package main

import (
	"fmt"
	"port_scanning/concurrent_scanning"
	"port_scanning/non_concurrent_scanning"
)

func main() {
	non_concurrent_scanning.ScanPorts()
	fmt.Println("-----------")
	concurrent_scanning.ScanPorts()
}
