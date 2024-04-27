package main

import (
	"fmt"

	"github.com/dimanchick22/service_controllers/discovery"
)

// ProxyService представляет сервис прокси
type ProxyService struct {
	DiscoveryService *discovery.DiscoveryService
}

// NewProxyService создает новый экземпляр сервиса прокси


// Здесь можно добавить методы для использования сервиса прокси

func main() {
	configFile := "config.txt" // Имя файла конфигурации

	// Создаем экземпляр сервиса обнаружения
	discoveryService := discovery.NewDiscoveryService(configFile)

	// Сканируем порты
	portMap := discoveryService.ScanPorts()
	if portMap == nil {
		fmt.Println("Ошибка сканирования портов")
		return
	}

}