package utils

import "os"

func ReadFile(fileName string) (string, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func WriteFile(fileName, data string) error {
	return os.WriteFile(fileName, []byte(data), 0644)
}
