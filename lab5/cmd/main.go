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

	C1List := []utils.Point{
		{X: big.NewInt(188), Y: big.NewInt(93)},
		{X: big.NewInt(725), Y: big.NewInt(195)},
		{X: big.NewInt(346), Y: big.NewInt(242)},
		{X: big.NewInt(489), Y: big.NewInt(468)},
		{X: big.NewInt(725), Y: big.NewInt(195)},
		{X: big.NewInt(745), Y: big.NewInt(210)},
		{X: big.NewInt(725), Y: big.NewInt(195)},
		{X: big.NewInt(618), Y: big.NewInt(206)},
		{X: big.NewInt(286), Y: big.NewInt(136)},
		{X: big.NewInt(179), Y: big.NewInt(275)},
	}

	C2List := []utils.Point{
		{X: big.NewInt(623), Y: big.NewInt(166)},
		{X: big.NewInt(513), Y: big.NewInt(414)},
		{X: big.NewInt(461), Y: big.NewInt(4)},
		{X: big.NewInt(739), Y: big.NewInt(574)},
		{X: big.NewInt(663), Y: big.NewInt(476)},
		{X: big.NewInt(724), Y: big.NewInt(522)},
		{X: big.NewInt(663), Y: big.NewInt(476)},
		{X: big.NewInt(438), Y: big.NewInt(40)},
		{X: big.NewInt(546), Y: big.NewInt(670)},
		{X: big.NewInt(482), Y: big.NewInt(230)},
	}

	PmList := utils.DecryptWord(C1List, C2List, nb)

	plaintext := utils.PointsToCharacters(PmList)

	fmt.Println("Дешифрованный текст:", plaintext)
}
