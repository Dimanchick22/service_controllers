package portscanner // Объявление пакета portscanner

import (
	"bufio"
	"os"
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

    scanner := bufio.NewScanner(file) // Создание сканера для файла
    for scanner.Scan() { 
        line := scanner.Text() // Чтение строки из файла
        fields := strings.Fields(line) // Разбиение строки на подстроки по пробелам
 
        host := fields[0] // Извлечение имени хоста из первой подстроки
        
        ports := make([]int, 0) // Создание пустого среза для хранения портов

        portMap[host] = ports // Запись среза портов в карту по ключу - имени хоста

    }

    if err := scanner.Err(); err != nil { // Ошибки при сканировании
        return nil, err 
    }

    return portMap, nil // Возврат всего
}
