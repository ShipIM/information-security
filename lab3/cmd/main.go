package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ShipIM/information-security/lab3/utils"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите N: ")
	NStr, _ := reader.ReadString('\n')
	NStr = strings.TrimSpace(NStr)

	fmt.Print("Введите e1: ")
	e1Str, _ := reader.ReadString('\n')
	e1Str = strings.TrimSpace(e1Str)

	fmt.Print("Введите e2: ")
	e2Str, _ := reader.ReadString('\n')
	e2Str = strings.TrimSpace(e2Str)

	fmt.Print("Введите путь к файлу C1: ")
	C1FilePath, _ := reader.ReadString('\n')
	C1FilePath = strings.TrimSpace(C1FilePath)

	fmt.Print("Введите путь к файлу C2: ")
	C2FilePath, _ := reader.ReadString('\n')
	C2FilePath = strings.TrimSpace(C2FilePath)

	C1Values, err := utils.ReadLinesFromFile(C1FilePath)
	if err != nil {
		fmt.Println("Ошибка при чтении из файла C1:", err)
		return
	}

	C2Values, err := utils.ReadLinesFromFile(C2FilePath)
	if err != nil {
		fmt.Println("Ошибка при чтении из файла C2:", err)
		return
	}

	N := new(big.Int)
	e1 := new(big.Int)
	e2 := new(big.Int)

	N.SetString(NStr, 10)
	e1.SetString(e1Str, 10)
	e2.SetString(e2Str, 10)

	_, r, s := utils.ExtendedGCD(e1, e2)

	fmt.Println("r =", r)
	fmt.Printf("s =%s\n", s)

	var resultBytes []byte
	for i := 0; i < len(C1Values) && i < len(C2Values); i++ {
		C1 := new(big.Int)
		C2 := new(big.Int)

		C1.SetString(C1Values[i], 10)
		C2.SetString(C2Values[i], 10)

		C1r := new(big.Int).Exp(C1, r, N)
		fmt.Printf("\nC1^r mod N для C1[%d] = %s: %s\n", i, C1Values[i], C1r)

		C2s := new(big.Int).Exp(C2, s, N)
		fmt.Printf("C2^s mod N для C2[%d] = %s: %s\n", i, C2Values[i], C2s)

		m := new(big.Int).Mod(new(big.Int).Mul(C1r, C2s), N)
		fmt.Printf("Значение m для C1[%d] и C2[%d] = %s\n", i, i, m.String())

		resultBytes = append(resultBytes, m.Bytes()...)
	}

	decodedText, err := utils.DecodeCP1251(resultBytes)
	if err != nil {
		fmt.Println("Ошибка при декодировании текста:", err)
	}

	fmt.Printf("\nДекодированный текст: %s", decodedText)
}
