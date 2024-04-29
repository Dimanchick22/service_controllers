package main

import (
	"flag"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/dimanchick22/service_controllers/discovery"
)

// ProxyService представляет сервис прокси
type ProxyService struct {
	DiscoveryService *discovery.DiscoveryService
}

var StartPorts = 8000
var EndPorts = 8050

// NewProxyService создает новый экземпляр сервиса прокси
func NewProxyService(discoveryService *discovery.DiscoveryService) *ProxyService {
	return &ProxyService{
		DiscoveryService: discoveryService,
	}
}

// SendRequest отправляет запрос через сервис прокси
func (ps *ProxyService) SendRequest(portMap map[string][]int) (string, error) {
	// Проверяем, есть ли доступные хосты и порты
	if len(portMap) == 0 {
		return "", fmt.Errorf("пустая мапа с портами")
	}

	// Получение случайного хоста
	host := getRandomHost(portMap)

	// Получение портов для выбранного хоста
	ports := portMap[host]

	// Проверяем, есть ли доступные порты для выбранного хоста
	if len(ports) == 0 {
		return "", fmt.Errorf("нет доступных портов для хоста %s", host)
	}

	// Выбор случайного порта
	port := ports[rand.Intn(len(ports))]

	// Формирование URL для запроса
	url := fmt.Sprintf("http://%s:%d", host, port)

	// Отправка GET-запроса на выбранный порт
	resp, err := http.Get(url)
	if err != nil {
		return "", fmt.Errorf("ошибка при выполнении запроса: %v", err)
	}
	defer resp.Body.Close()

	// Возвращаем ответ от сервера
	return fmt.Sprintf("Ответ от сервера (%s): %s\n", url, resp.Status), nil
}

// getRandomHost выбирает случайный хост из мапы с портами
func getRandomHost(portMap map[string][]int) string {
	var hosts []string
	for host := range portMap {
		hosts = append(hosts, host)
	}
	rand.Seed(time.Now().UnixNano())
	return hosts[rand.Intn(len(hosts))]
}

// getMissingPort возвращает один порт в указанном диапазоне, который отсутствует в переданной мапе портов
func getMissingPort(portMap map[string][]int, startPort, endPort int) (int, error) {
	// Проверка на неверно указанный диапазон портов
	if startPort > endPort {
		return 0, fmt.Errorf("неверно указан диапазон портов: startPort больше endPort")
	}

	// Создаем карту для быстрой проверки существующих портов
	existingPorts := make(map[int]bool)
	for _, ports := range portMap {
		for _, port := range ports {
			existingPorts[port] = true
		}
	}

	// Проверяем каждый порт в диапазоне
	for port := startPort; port <= endPort; port++ {
		// Если порт отсутствует в мапе, возвращаем его
		if !existingPorts[port] {
			return port, nil
		}
	}

	// Если все порты в диапазоне уже заняты, возвращаем ошибку
	return 0, fmt.Errorf("в указанном диапазоне нет доступных портов")
}

// fillPortMap заполняет карту портов с использованием сервиса обнаружения
func fillPortMap(discoveryService *discovery.DiscoveryService) map[string][]int {
	portMap := make(map[string][]int)

	// Получаем карту портов с помощью сервиса обнаружения
	discoveredPorts := discoveryService.ScanPorts()

	// Копируем найденные порты в нашу карту
	for host, ports := range discoveredPorts {
		portMap[host] = ports
	}

	return portMap
}

// sendRequestHandler обрабатывает запрос на эндпоинте /send-request
func sendRequestHandler(w http.ResponseWriter, r *http.Request, ps *ProxyService) {
	// Отправляем запрос через сервис прокси
	portMap := ps.DiscoveryService.ScanPorts() // Сначала сканируем порты
	if portMap == nil {
		http.Error(w, "Ошибка получения мапы с портами", http.StatusInternalServerError)
		return
	}
	response, err := ps.SendRequest(portMap)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка отправки запроса: %v", err), http.StatusInternalServerError)
		return
	}

	// Отправляем ответ пользователю
	fmt.Fprint(w, response)
}

// portHandler обрабатывает запрос на эндпоинте /port
func portHandler(w http.ResponseWriter, r *http.Request, ps *ProxyService, discoveryService *discovery.DiscoveryService, configFile string) {
	// Заполняем карту портов с помощью сервиса обнаружения
	portMap := fillPortMap(discoveryService)

	// Получаем один отсутствующий порт в заданном диапазоне
	missingPort, err := getMissingPort(portMap, StartPorts, EndPorts)
	if err != nil {
		http.Error(w, fmt.Sprintf("Ошибка получения отсутствующего порта: %v", err), http.StatusInternalServerError)
		return
	}

	// Отправляем отсутствующий порт пользователю
	fmt.Fprintf(w, "Отсутствующий порт: %d\n", missingPort)
}

// main - точка входа в программу
func main() {
	// Обработка флага -f для указания конфигурационного файла
	configFile := flag.String("f", "config.txt", "Путь к файлу конфигурации")
	flag.Parse()


	// Создаем экземпляр сервиса обнаружения
	discoveryService := discovery.NewDiscoveryService(*configFile,StartPorts,EndPorts)

	portMap := discoveryService.ScanPorts()
	if portMap == nil {
		fmt.Println("Ошибка сканирования портов. Завершение программы.")
		return
	}
	// Создаем экземпляр сервиса прокси
	proxyService := NewProxyService(discoveryService)

	// Настройка обработчика для эндпоинта /send-request
	http.HandleFunc("/send-request", func(w http.ResponseWriter, r *http.Request) {
		sendRequestHandler(w, r, proxyService)
	})

	// Настройка обработчика для эндпоинта /port
	http.HandleFunc("/port", func(w http.ResponseWriter, r *http.Request) {
		portHandler(w, r, proxyService, discoveryService, *configFile)
	})

	// Запуск веб-сервера на порту 9000
	fmt.Println("Сервер запущен на порту 9000")
	if err := http.ListenAndServe(":9000", nil); err != nil {
		fmt.Printf("Ошибка запуска веб-сервера: %v\n", err)
	}
}
