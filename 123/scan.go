package scan

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func scanPort(host string, port int, wg *sync.WaitGroup) {
	defer wg.Done()
	target := fmt.Sprintf("%s:%d", host, port)
	conn, err := net.DialTimeout("tcp", target, 500*time.Millisecond)
	if err == nil {
			fmt.Printf("Port %d: Open\n", port)
			conn.Close()
	}
}

func main() {
	host := "example.com"
	startPort := 1
	endPort := 65535
	var wg sync.WaitGroup

	fmt.Printf("Scanning ports on %s...\n", host)

	for port := startPort; port <= endPort; port++ {
			wg.Add(1)
			go scanPort(host, port, &wg)
	}

	wg.Wait()
}