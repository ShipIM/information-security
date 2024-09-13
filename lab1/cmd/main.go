package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ShipIM/information-security/lab1/utils"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите имя файла для шифрования/дешифрации: ")
	fileName, _ := reader.ReadString('\n')
	fileName = strings.TrimSpace(fileName)

	plainText, err := utils.ReadFile(fileName)
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}

	fmt.Println("Введите ключи (через пробел):")
	keysInput, _ := reader.ReadString('\n')
	keys := strings.Fields(keysInput)

	composedKey := utils.ComposeKey(keys)
	fmt.Println("Составной ключ:", composedKey)

	fmt.Print("Выберите режим (encrypt/decrypt): ")
	mode, _ := reader.ReadString('\n')
	mode = strings.TrimSpace(mode)

	var result string
	switch mode {
	case "encrypt":
		result = utils.VigenereEncrypt(plainText, composedKey)
	case "decrypt":
		result = utils.VigenereDecrypt(plainText, composedKey)
	default:
		fmt.Println("Неправильный режим работы")
		os.Exit(1)
	}

	outputFileName := "output.txt"
	if err := utils.WriteFile(outputFileName, result); err != nil {
		fmt.Println("Ошибка записи в файл:", err)
	} else {
		fmt.Println("Результат сохранён в файл:", outputFileName)
	}
}
