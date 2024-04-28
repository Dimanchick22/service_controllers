package portscanner // Объявление пакета portscanner

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// ReadConfig читает файл конфигурации и возвращает мапу хостов и диапазона портов
func ReadConfig(filename string) (map[string][]int, error) { 

    file, err := os.Open(filename) 
    if err != nil { 
        return nil,  err 
    }
    defer file.Close() // Отложенное закрытие файла при выходе из функции

    portMap := make(map[string][]int) // Создание карты для хранения хостов и их портов
    var startPort, endPort int // Объявление переменных startPort - хранит начальный порт, endPort - хранит конечный порт

    scanner := bufio.NewScanner(file) // Создание сканера для файла
    for scanner.Scan() { 
        line := scanner.Text() // Чтение строки из файла
        fields := strings.Fields(line) // Разбиение строки на подстроки по пробелам
        if len(fields) < 3 { // Проверка количества подстрок
            continue // Пропуск строки
        }
        host := fields[0] // Извлечение имени хоста из первой подстроки
        start, err := strconv.Atoi(fields[1]) // Преобразование второй подстроки в число
        if err != nil { 
            continue 
        }
        end, err := strconv.Atoi(fields[2]) // Преобразование третьей подстроки в число
        if err != nil { 
            continue 
        }
        ports := make([]int, 0) // Создание пустого среза для хранения портов
        for port := start; port <= end; port++ { // Цикл по всем портам в заданном диапазоне
            ports = append(ports, port) // Добавление порта в срез портов
        }
        portMap[host] = ports // Запись среза портов в карту по ключу - имени хоста
        if startPort == 0 || start < startPort { // Проверка и обновление начального порта
            startPort = start
        }
        if endPort == 0 || end > endPort { // Проверка и обновление конечного порта
            endPort = end
        }
    }

    if err := scanner.Err(); err != nil { // Ошибки при сканировании
        return nil, err 
    }

    return portMap, nil // Возврат всего
}
