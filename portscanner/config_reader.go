package portscanner

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

// ReadConfig читает файл конфигурации и возвращает мапу хостов и диапазона портов
func ReadConfig(filename string) (map[string][]int, int, int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, 0, 0, err
	}
	defer file.Close()

	portMap := make(map[string][]int)
	var startPort, endPort int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fields) < 3 {
			continue
		}
		host := fields[0]
		start, err := strconv.Atoi(fields[1])
		if err != nil {
			continue
		}
		end, err := strconv.Atoi(fields[2])
		if err != nil {
			continue
		}
		ports := make([]int, 0)
		for port := start; port <= end; port++ {
			ports = append(ports, port)
		}
		portMap[host] = ports

		if startPort == 0 || start < startPort {
			startPort = start
		}
		if endPort == 0 || end > endPort {
			endPort = end
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, 0, 0, err
	}

	return portMap, startPort, endPort, nil
}
