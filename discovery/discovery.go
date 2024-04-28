package discovery

import (
	"fmt"

	"github.com/dimanchick22/service_controllers/portscanner"
)

// DiscoveryService предоставляет методы для обнаружения доступных сервисов и портов
type DiscoveryService struct {
	ConfigFile string
	PortMap    map[string][]int
	StartPort  int
	EndPort    int
}

// NewDiscoveryService создает новый DiscoveryService с указанным файлом конфигурации
func NewDiscoveryService(configFile string, startPort, endPort int) *DiscoveryService {
	return &DiscoveryService{
		ConfigFile: configFile,
		StartPort:  startPort,
		EndPort:    endPort,
	}
}

// ScanPorts сканирует порты для всех хостов в файле конфигурации и возвращает мапу с портами
func (ds *DiscoveryService) ScanPorts() map[string][]int {
	portMap, err := portscanner.ReadConfig(ds.ConfigFile)
	if err != nil {
		fmt.Printf("Ошибка чтения файла конфигурации: %v\n", err)
		return nil
	}
	ds.PortMap = portMap

	results := make(map[string][]int)

	for host := range portMap {
		portscanner.ScanPorts(host, ds.StartPort, ds.EndPort, results)
	}

	return results
}
func (ds *DiscoveryService) GetPortMap() map[string][]int {
	return ds.PortMap
}