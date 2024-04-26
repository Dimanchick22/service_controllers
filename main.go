package main

import (
	"fmt"
	"net"
	"time"
)

// Функция сканирования портов на заданном хосте в указанном диапазоне и обновления мапы с найденными портами
func scanPorts(host string, startPort, endPort int, portMap map[string][]int) {
	fmt.Printf("Начало сканирования портов на хосте %s...\n", host)
	for port := startPort; port <= endPort; port++ {
		target := fmt.Sprintf("%s:%d", host, port)
		_, err := net.DialTimeout("tcp", target, 500*time.Millisecond)
		if err == nil {
			portMap[host] = append(portMap[host], port)
			fmt.Printf("Порт %d на хосте %s обнаружен и добавлен в мапу\n", port, host)
		}
	}
	fmt.Printf("Завершение сканирования портов на хосте %s\n", host)
}

func main() {
	// Заданные хосты и диапазон портов для сканирования
	hosts := []string{"127.0.0.1", "192.168.1.11"} // Замените на ваши хосты
	startPort := 8000
	endPort := 8200

	// Мапа для хранения найденных портов
	portMap := make(map[string][]int)

	// Бесконечный цикл для постоянного сканирования портов
	for {
		fmt.Println("Начало нового цикла сканирования портов...")

		// Сканирование портов для каждого хоста
		for _, host := range hosts {
			// Сканирование портов и обновление мапы
			scanPorts(host, startPort, endPort, portMap)
		}

		// Вывод актуальной информации о портах
		fmt.Println("Актуальная информация о портах:")
		for host, ports := range portMap {
			fmt.Printf("Хост: %s, Порты: %v\n", host, ports)
		}

		fmt.Println("Ожидание перед следующим сканированием...")
		// Пауза перед следующим сканированием
		time.Sleep(20 * time.Second)
	}
}
