package utils

import (
	"bufio"
	"os"
	"strings"

	"golang.org/x/text/encoding/charmap"
)

func ReadLinesFromFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, strings.TrimSpace(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

func DecodeCP1251(b []byte) (string, error) {
	decoded, err := charmap.Windows1251.NewDecoder().Bytes(b)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}
