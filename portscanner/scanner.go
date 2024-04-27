package portscanner

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// ScanPorts сканирует порты на заданном хосте в указанном диапазоне и обновляет мапу с найденными портами
func ScanPorts(host string, startPort, endPort int, portMap map[string][]int) {
	fmt.Printf("Начало сканирования портов на хосте %s...\n", host)

	// Очистка мапы перед сканированием
	portMap[host] = []int{}

	var wg sync.WaitGroup
	var mu sync.Mutex

	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		go func(p int) {
			defer wg.Done()
			target := fmt.Sprintf("%s:%d", host, p)
			conn, err := net.DialTimeout("tcp", target, 500*time.Millisecond)
			if err == nil {
				conn.Close()
				mu.Lock()
				portMap[host] = append(portMap[host], p)
				mu.Unlock()
				fmt.Printf("Порт %d на хосте %s обнаружен и добавлен в мапу\n", p, host)
			}
		}(port)
	}

	wg.Wait()
	fmt.Printf("Завершение сканирования портов на хосте %s\n", host)
}
