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

	fmt.Print("Выберите режим (encrypt/decrypt): ")
	mode, _ := reader.ReadString('\n')
	mode = strings.TrimSpace(mode)

	var result string

	switch mode {
	case "encrypt":
		result, err = utils.VigenereEncrypt(plainText, keys)
		if err != nil {
			fmt.Println("Ошибка при шифровании:", err)
			return
		}
	case "decrypt":
		result, err = utils.VigenereDecrypt(plainText, keys)
		if err != nil {
			fmt.Println("Ошибка при дешифровании:", err)
			return
		}
	default:
		fmt.Println("Неправильный режим работы")
		os.Exit(1)
	}

	outputFileName := "files/output.txt"
	if err := utils.WriteFile(outputFileName, result); err != nil {
		fmt.Println("Ошибка записи в файл:", err)
	} else {
		fmt.Println("Результат сохранён в файл:", outputFileName)
	}
}
