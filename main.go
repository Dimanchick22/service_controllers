package main

import (
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

// NewProxyService создает новый экземпляр сервиса прокси
func NewProxyService(discoveryService *discovery.DiscoveryService) *ProxyService {
	return &ProxyService{
		DiscoveryService: discoveryService,
	}
}

// SendRequest отправляет запрос на случайный порт хоста из мапы
func (ps *ProxyService) SendRequest(portMap map[string][]int) (string, error) {
	// Получение случайного хоста
	host := getRandomHost(portMap)

	// Получение портов для выбранного хоста
	ports := portMap[host]

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

// Функция для выбора случайного хоста из мапы с портами
func getRandomHost(portMap map[string][]int) string {
	var hosts []string
	for host := range portMap {
		hosts = append(hosts, host)
	}
	rand.Seed(time.Now().UnixNano())
	return hosts[rand.Intn(len(hosts))]
}

// Обработчик для эндпоинта /send-request
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

func main() {
	configFile := "config.txt" // Имя файла конфигурации

	// Создаем экземпляр сервиса обнаружения
	discoveryService := discovery.NewDiscoveryService(configFile)

	// Создаем экземпляр сервиса прокси
	proxyService := NewProxyService(discoveryService)

	// Настройка обработчика для эндпоинта /send-request
	http.HandleFunc("/send-request", func(w http.ResponseWriter, r *http.Request) {
		sendRequestHandler(w, r, proxyService)
	})

	// Запуск веб-сервера на порту 8080
	fmt.Println("Сервер запущен на порту 9000")
	http.ListenAndServe(":9000", nil)
}
