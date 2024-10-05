package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ShipIM/information-security/lab2/utils"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите N: ")
	NStr, _ := reader.ReadString('\n')
	NStr = strings.TrimSpace(NStr)

	fmt.Print("Введите e: ")
	e1Str, _ := reader.ReadString('\n')
	e1Str = strings.TrimSpace(e1Str)

	fmt.Print("Введите путь к файлу C: ")
	CFilePath, _ := reader.ReadString('\n')
	CFilePath = strings.TrimSpace(CFilePath)

	CValues, err := utils.ReadLinesFromFile(CFilePath)
	if err != nil {
		fmt.Println("Ошибка при чтении из файла C:", err)
		return
	}

	N := new(big.Int)
	e := new(big.Int)

	N.SetString(NStr, 10)
	e.SetString(e1Str, 10)

	var resultBytes []byte
	for i := 0; i < len(CValues); i++ {
		C := new(big.Int)

		C.SetString(CValues[i], 10)

		res := utils.Encode(N, e, C)

		resultBytes = append(resultBytes, res.Bytes()...)
	}

	decodedText, err := utils.DecodeCP1251(resultBytes)
	if err != nil {
		fmt.Println("Ошибка при декодировании текста:", err)
	}

	fmt.Printf("\nДекодированный текст: %s", decodedText)
}
