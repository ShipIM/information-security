package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"

	"github.com/ShipIM/information-security/lab4/utils"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите открытый ключ (формат: x,y):")
	publicKeyInput, _ := reader.ReadString('\n')
	publicKeyInput = strings.TrimSpace(publicKeyInput)

	var x, y big.Int
	fmt.Sscanf(publicKeyInput, "%d,%d", &x, &y)
	Pb := utils.Point{X: &x, Y: &y}

	fmt.Println("Введите слово в формате: A1 B2 C3 D4 ...")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	word, kList := utils.ParseInput(input)

	C1List, C2List := utils.EncryptWord(word, kList, Pb)

	fmt.Println("Шифрованный текст:")
	for i := range word {
		fmt.Printf("C1: (%s, %s), C2: (%s, %s) для буквы %c\n",
			C1List[i].X.String(), C1List[i].Y.String(),
			C2List[i].X.String(), C2List[i].Y.String(),
			word[i],
		)
	}
}
