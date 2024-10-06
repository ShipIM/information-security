package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ShipIM/information-security/lab5/utils"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите секретное число (nb):")
	nbInput, _ := reader.ReadString('\n')
	nbInput = strings.TrimSpace(nbInput)
	nb := new(big.Int)
	nb.SetString(nbInput, 10)

	fmt.Print("Введите путь к файлу C1: ")
	C1FilePath, _ := reader.ReadString('\n')
	C1FilePath = strings.TrimSpace(C1FilePath)

	fmt.Print("Введите путь к файлу C2: ")
	C2FilePath, _ := reader.ReadString('\n')
	C2FilePath = strings.TrimSpace(C2FilePath)

	C1List, err := utils.ReadPointsFromFile(C1FilePath)
	if err != nil {
		fmt.Println("Ошибка при чтении из файла C1:", err)
		return
	}

	C2List, err := utils.ReadPointsFromFile(C2FilePath)
	if err != nil {
		fmt.Println("Ошибка при чтении из файла C2:", err)
		return
	}

	PmList := utils.DecryptWord(C1List, C2List, nb)

	plaintext := utils.PointsToCharacters(PmList)

	fmt.Println("Дешифрованный текст:", plaintext)
}
