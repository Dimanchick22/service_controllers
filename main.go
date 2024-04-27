package main

import (
	"fmt"
	"time"

	"github.com/dimanchick22/service_controllers/portscanner"
)

func main() {
	configFile := "config.txt" // Имя файла конфигурации
	// Вызов новой функции для сканирования портов
	scanPortsForever(configFile)
}

// Функция для сканирования портов и вывода результатов в бесконечном цикле
func scanPortsForever(configFile string) {
	// Чтение конфигурационного файла
	portMap, startPort, endPort, err := portscanner.ReadConfig(configFile)
	if err != nil {
		fmt.Printf("Ошибка чтения файла конфигурации: %v\n", err)
		return
	}

	// Мапа для хранения найденных портов
	results := make(map[string][]int)

	// Бесконечный цикл для постоянного сканирования портов
	for {
		fmt.Println("Начало нового цикла сканирования портов...")

		// Сканирование портов для каждого хоста
		for host:= range portMap {
			// Сканирование портов и обновление мапы
			portscanner.ScanPorts(host, startPort, endPort, results)
		}

		// Вывод актуальной информации о портах
		fmt.Println("Актуальная информация о портах:")
		for host, ports := range results {
			fmt.Printf("Хост: %s, Порты: %v\n", host, ports)
		}

		fmt.Println("Ожидание перед следующим сканированием...")
		// Пауза перед следующим сканированием
		time.Sleep(20 * time.Second)
	}
}
