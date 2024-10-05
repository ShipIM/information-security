package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ShipIM/information-security/lab3/utils"
	"golang.org/x/text/encoding/charmap"
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

	fmt.Print("Введите C1: ")
	C1Str, _ := reader.ReadString('\n')
	C1Str = strings.TrimSpace(C1Str)

	fmt.Print("Введите C2: ")
	C2Str, _ := reader.ReadString('\n')
	C2Str = strings.TrimSpace(C2Str)

	N := new(big.Int)
	e1 := new(big.Int)
	e2 := new(big.Int)
	C1 := new(big.Int)
	C2 := new(big.Int)

	N.SetString(NStr, 10)
	e1.SetString(e1Str, 10)
	e2.SetString(e2Str, 10)
	C1.SetString(C1Str, 10)
	C2.SetString(C2Str, 10)

	_, r, s := utils.ExtendedGCD(e1, e2)

	fmt.Println("r =", r)
	fmt.Println("s =", s)

	C1r := new(big.Int).Exp(C1, r, N)
	fmt.Println("C1^r mod N =", C1r)

	C2s := new(big.Int).Exp(C2, s, N)
	fmt.Println("C2^s mod N =", C2s)

	m := new(big.Int).Mod(new(big.Int).Mul(C1r, C2s), N)
	fmt.Println("m mod N =", m)

	fmt.Printf("Значение m = %s\n", m.String())

	mBytes := m.Bytes()
	decodedText, err := decodeCP1251(mBytes)
	if err != nil {
		fmt.Println("Ошибка при декодировании текста:", err)
		return
	}

	fmt.Printf("Декодированный текст: %s\n", decodedText)
}

func decodeCP1251(b []byte) (string, error) {
	decoded, err := charmap.Windows1251.NewDecoder().Bytes(b)
	if err != nil {
		return "", err
	}

	return string(decoded), nil
}
