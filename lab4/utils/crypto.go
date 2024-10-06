package utils

import (
	"fmt"
	"math/big"
	"strings"
)

type Point struct {
	X, Y *big.Int
}

// Параметры эллиптической кривой E751: y^2 = x^3 - x + 1 (mod 751)
var (
	a = big.NewInt(-1)                      // Коэффициент a
	p = big.NewInt(751)                     // Модуль кривой
	G = Point{big.NewInt(0), big.NewInt(1)} // Генерирующая точка G = (0, 1)
)

var alphabet = map[rune]Point{
	' ':  {big.NewInt(33), big.NewInt(355)},
	'!':  {big.NewInt(33), big.NewInt(396)},
	'"':  {big.NewInt(34), big.NewInt(74)},
	'#':  {big.NewInt(34), big.NewInt(677)},
	'$':  {big.NewInt(36), big.NewInt(87)},
	'%':  {big.NewInt(36), big.NewInt(664)},
	'&':  {big.NewInt(39), big.NewInt(171)},
	'\'': {big.NewInt(39), big.NewInt(580)},
	'(':  {big.NewInt(43), big.NewInt(224)},
	')':  {big.NewInt(43), big.NewInt(527)},
	'*':  {big.NewInt(44), big.NewInt(366)},
	'+':  {big.NewInt(44), big.NewInt(385)},
	',':  {big.NewInt(45), big.NewInt(31)},
	'-':  {big.NewInt(45), big.NewInt(720)},
	'.':  {big.NewInt(47), big.NewInt(349)},
	'/':  {big.NewInt(47), big.NewInt(402)},
	'0':  {big.NewInt(48), big.NewInt(49)},
	'1':  {big.NewInt(48), big.NewInt(702)},
	'2':  {big.NewInt(49), big.NewInt(183)},
	'3':  {big.NewInt(49), big.NewInt(568)},
	'4':  {big.NewInt(53), big.NewInt(277)},
	'5':  {big.NewInt(53), big.NewInt(474)},
	'6':  {big.NewInt(56), big.NewInt(332)},
	'7':  {big.NewInt(56), big.NewInt(419)},
	'8':  {big.NewInt(58), big.NewInt(139)},
	'9':  {big.NewInt(58), big.NewInt(612)},
	':':  {big.NewInt(59), big.NewInt(365)},
	';':  {big.NewInt(59), big.NewInt(386)},
	'<':  {big.NewInt(61), big.NewInt(129)},
	'=':  {big.NewInt(61), big.NewInt(622)},
	'>':  {big.NewInt(62), big.NewInt(372)},
	'?':  {big.NewInt(62), big.NewInt(379)},
	'@':  {big.NewInt(66), big.NewInt(199)},
	'A':  {big.NewInt(66), big.NewInt(552)},
	'B':  {big.NewInt(67), big.NewInt(84)},
	'C':  {big.NewInt(67), big.NewInt(667)},
	'D':  {big.NewInt(69), big.NewInt(241)},
	'E':  {big.NewInt(69), big.NewInt(510)},
	'F':  {big.NewInt(70), big.NewInt(195)},
	'G':  {big.NewInt(70), big.NewInt(556)},
	'H':  {big.NewInt(72), big.NewInt(254)},
	'I':  {big.NewInt(72), big.NewInt(497)},
	'J':  {big.NewInt(73), big.NewInt(72)},
	'K':  {big.NewInt(73), big.NewInt(679)},
	'L':  {big.NewInt(74), big.NewInt(170)},
	'M':  {big.NewInt(74), big.NewInt(581)},
	'N':  {big.NewInt(75), big.NewInt(318)},
	'O':  {big.NewInt(75), big.NewInt(433)},
	'P':  {big.NewInt(78), big.NewInt(271)},
	'Q':  {big.NewInt(78), big.NewInt(480)},
	'R':  {big.NewInt(79), big.NewInt(111)},
	'S':  {big.NewInt(79), big.NewInt(640)},
	'T':  {big.NewInt(80), big.NewInt(318)},
	'U':  {big.NewInt(80), big.NewInt(433)},
	'V':  {big.NewInt(82), big.NewInt(270)},
	'W':  {big.NewInt(82), big.NewInt(481)},
	'X':  {big.NewInt(83), big.NewInt(373)},
	'Y':  {big.NewInt(83), big.NewInt(378)},
	'Z':  {big.NewInt(85), big.NewInt(35)},
	'[':  {big.NewInt(85), big.NewInt(716)},
	'\\': {big.NewInt(86), big.NewInt(25)},
	']':  {big.NewInt(86), big.NewInt(726)},
	'^':  {big.NewInt(90), big.NewInt(21)},
	'_':  {big.NewInt(90), big.NewInt(730)},
	'`':  {big.NewInt(93), big.NewInt(267)},
	'a':  {big.NewInt(93), big.NewInt(484)},
	'b':  {big.NewInt(98), big.NewInt(338)},
	'c':  {big.NewInt(98), big.NewInt(413)},
	'd':  {big.NewInt(99), big.NewInt(295)},
	'e':  {big.NewInt(99), big.NewInt(456)},
	'f':  {big.NewInt(100), big.NewInt(364)},
	'g':  {big.NewInt(100), big.NewInt(387)},
	'h':  {big.NewInt(102), big.NewInt(267)},
	'i':  {big.NewInt(102), big.NewInt(484)},
	'j':  {big.NewInt(105), big.NewInt(369)},
	'k':  {big.NewInt(105), big.NewInt(382)},
	'l':  {big.NewInt(106), big.NewInt(24)},
	'm':  {big.NewInt(106), big.NewInt(727)},
	'n':  {big.NewInt(108), big.NewInt(247)},
	'o':  {big.NewInt(108), big.NewInt(504)},
	'p':  {big.NewInt(109), big.NewInt(200)},
	'q':  {big.NewInt(109), big.NewInt(551)},
	'r':  {big.NewInt(110), big.NewInt(129)},
	's':  {big.NewInt(110), big.NewInt(622)},
	't':  {big.NewInt(114), big.NewInt(144)},
	'u':  {big.NewInt(114), big.NewInt(607)},
	'v':  {big.NewInt(115), big.NewInt(242)},
	'w':  {big.NewInt(115), big.NewInt(509)},
	'x':  {big.NewInt(116), big.NewInt(92)},
	'y':  {big.NewInt(116), big.NewInt(659)},
	'z':  {big.NewInt(120), big.NewInt(147)},
	'{':  {big.NewInt(120), big.NewInt(604)},
	'|':  {big.NewInt(125), big.NewInt(292)},
	'}':  {big.NewInt(125), big.NewInt(459)},
	'~':  {big.NewInt(126), big.NewInt(33)},
	'А':  {big.NewInt(189), big.NewInt(297)},
	'Б':  {big.NewInt(189), big.NewInt(454)},
	'В':  {big.NewInt(192), big.NewInt(32)},
	'Г':  {big.NewInt(192), big.NewInt(719)},
	'Д':  {big.NewInt(194), big.NewInt(205)},
	'Е':  {big.NewInt(194), big.NewInt(546)},
	'Ж':  {big.NewInt(197), big.NewInt(145)},
	'З':  {big.NewInt(197), big.NewInt(606)},
	'И':  {big.NewInt(198), big.NewInt(224)},
	'Й':  {big.NewInt(198), big.NewInt(527)},
	'К':  {big.NewInt(200), big.NewInt(30)},
	'Л':  {big.NewInt(200), big.NewInt(721)},
	'М':  {big.NewInt(203), big.NewInt(324)},
	'Н':  {big.NewInt(203), big.NewInt(427)},
	'О':  {big.NewInt(205), big.NewInt(372)},
	'П':  {big.NewInt(205), big.NewInt(379)},
	'Р':  {big.NewInt(206), big.NewInt(106)},
	'С':  {big.NewInt(206), big.NewInt(645)},
	'Т':  {big.NewInt(209), big.NewInt(82)},
	'У':  {big.NewInt(209), big.NewInt(669)},
	'Ф':  {big.NewInt(210), big.NewInt(31)},
	'Х':  {big.NewInt(210), big.NewInt(720)},
	'Ц':  {big.NewInt(215), big.NewInt(247)},
	'Ч':  {big.NewInt(215), big.NewInt(504)},
	'Ш':  {big.NewInt(218), big.NewInt(150)},
	'Щ':  {big.NewInt(218), big.NewInt(601)},
	'Ъ':  {big.NewInt(221), big.NewInt(138)},
	'Ы':  {big.NewInt(221), big.NewInt(613)},
	'Ь':  {big.NewInt(226), big.NewInt(9)},
	'Э':  {big.NewInt(226), big.NewInt(742)},
	'Ю':  {big.NewInt(227), big.NewInt(299)},
	'Я':  {big.NewInt(227), big.NewInt(452)},
	'а':  {big.NewInt(228), big.NewInt(271)},
	'б':  {big.NewInt(228), big.NewInt(480)},
	'в':  {big.NewInt(229), big.NewInt(151)},
	'г':  {big.NewInt(229), big.NewInt(600)},
	'д':  {big.NewInt(234), big.NewInt(164)},
	'е':  {big.NewInt(234), big.NewInt(587)},
	'ж':  {big.NewInt(235), big.NewInt(19)},
	'з':  {big.NewInt(235), big.NewInt(732)},
	'и':  {big.NewInt(236), big.NewInt(39)},
	'й':  {big.NewInt(236), big.NewInt(712)},
	'к':  {big.NewInt(237), big.NewInt(297)},
	'л':  {big.NewInt(237), big.NewInt(454)},
	'м':  {big.NewInt(238), big.NewInt(175)},
	'н':  {big.NewInt(238), big.NewInt(576)},
	'о':  {big.NewInt(240), big.NewInt(309)},
	'п':  {big.NewInt(240), big.NewInt(442)},
	'р':  {big.NewInt(243), big.NewInt(87)},
	'с':  {big.NewInt(243), big.NewInt(664)},
	'т':  {big.NewInt(247), big.NewInt(266)},
	'у':  {big.NewInt(247), big.NewInt(485)},
	'ф':  {big.NewInt(249), big.NewInt(183)},
	'х':  {big.NewInt(249), big.NewInt(568)},
	'ц':  {big.NewInt(250), big.NewInt(14)},
	'ч':  {big.NewInt(250), big.NewInt(737)},
	'ш':  {big.NewInt(251), big.NewInt(245)},
	'щ':  {big.NewInt(251), big.NewInt(506)},
	'ъ':  {big.NewInt(253), big.NewInt(211)},
	'ы':  {big.NewInt(253), big.NewInt(540)},
	'ь':  {big.NewInt(256), big.NewInt(121)},
	'э':  {big.NewInt(256), big.NewInt(630)},
	'ю':  {big.NewInt(257), big.NewInt(293)},
	'я':  {big.NewInt(257), big.NewInt(458)},
}

func addPoints(p1, p2 Point) Point {
	if p1.X == nil {
		return p2
	}
	if p2.X == nil {
		return p1
	}

	var lambda *big.Int
	if p1.X.Cmp(p2.X) == 0 && p1.Y.Cmp(p2.Y) == 0 {
		// lambda = (3*x1^2 + a) / (2*y1) mod p
		lambda = new(big.Int).Mul(big.NewInt(3), new(big.Int).Mul(p1.X, p1.X))
		lambda.Add(lambda, a)
		y2Inv := new(big.Int).ModInverse(new(big.Int).Mul(big.NewInt(2), p1.Y), p)
		lambda.Mul(lambda, y2Inv)
	} else {
		// lambda = (y2 - y1) / (x2 - x1) mod p
		lambda = new(big.Int).Sub(p2.Y, p1.Y)
		xDiff := new(big.Int).Sub(p2.X, p1.X)
		xDiffInv := new(big.Int).ModInverse(xDiff, p)
		lambda.Mul(lambda, xDiffInv)
	}
	lambda.Mod(lambda, p)

	// x3 = lambda^2 - x1 - x2 mod p
	x3 := new(big.Int).Mul(lambda, lambda)
	x3.Sub(x3, p1.X)
	x3.Sub(x3, p2.X)
	x3.Mod(x3, p)

	// y3 = lambda(x1 - x3) - y1 mod p
	y3 := new(big.Int).Sub(p1.X, x3)
	y3.Mul(lambda, y3)
	y3.Sub(y3, p1.Y)
	y3.Mod(y3, p)

	return Point{x3, y3}
}

func scalarMult(k *big.Int, p Point) Point {
	res := Point{nil, nil}
	tmp := p

	for i := k.BitLen() - 1; i >= 0; i-- {
		res = addPoints(res, res)
		if k.Bit(i) == 1 {
			res = addPoints(res, tmp)
		}
	}

	return res
}

func encryptPoint(Pm Point, k *big.Int, Pb Point) (Point, Point) {
	C1 := scalarMult(k, G)   // C1 = k * G
	kPb := scalarMult(k, Pb) // k * Pb
	C2 := addPoints(Pm, kPb) // C2 = Pm + k * Pb

	return C1, C2
}

func EncryptWord(word []rune, kList []*big.Int, Pb Point) ([]Point, []Point) {
	C1List := []Point{}
	C2List := []Point{}

	for i, letter := range word {
		Pm, exists := alphabet[letter]
		if !exists {
			continue
		}
		C1, C2 := encryptPoint(Pm, kList[i], Pb)
		C1List = append(C1List, C1)
		C2List = append(C2List, C2)
	}

	return C1List, C2List
}

func ParseInput(input string) ([]rune, []*big.Int) {
	parts := strings.Fields(input)
	word := make([]rune, len(parts))
	kList := make([]*big.Int, len(parts))

	for i, part := range parts {
		var letter rune
		var k big.Int
		fmt.Sscanf(part, "%c%d", &letter, &k)
		word[i] = letter
		kList[i] = new(big.Int).Set(&k)
	}

	return word, kList
}
