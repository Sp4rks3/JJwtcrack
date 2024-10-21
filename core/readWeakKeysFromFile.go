package core

import (
	"bufio"
	"os"
	"strings"
)

func readWeakKeysFromFile(filePath string) ([]string, error) {
	var weakKeys []string

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		weakKeys = append(weakKeys, strings.TrimSpace(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return weakKeys, nil
}
