package utils

import (
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

var alphabet = map[Point]rune{
	{big.NewInt(33), big.NewInt(355)}:  ' ',
	{big.NewInt(33), big.NewInt(396)}:  '!',
	{big.NewInt(34), big.NewInt(74)}:   '"',
	{big.NewInt(34), big.NewInt(677)}:  '#',
	{big.NewInt(36), big.NewInt(87)}:   '$',
	{big.NewInt(36), big.NewInt(664)}:  '%',
	{big.NewInt(39), big.NewInt(171)}:  '&',
	{big.NewInt(39), big.NewInt(580)}:  '\'',
	{big.NewInt(43), big.NewInt(224)}:  '(',
	{big.NewInt(43), big.NewInt(527)}:  ')',
	{big.NewInt(44), big.NewInt(366)}:  '*',
	{big.NewInt(44), big.NewInt(385)}:  '+',
	{big.NewInt(45), big.NewInt(31)}:   ',',
	{big.NewInt(45), big.NewInt(720)}:  '-',
	{big.NewInt(47), big.NewInt(349)}:  '.',
	{big.NewInt(47), big.NewInt(402)}:  '/',
	{big.NewInt(48), big.NewInt(49)}:   '0',
	{big.NewInt(48), big.NewInt(702)}:  '1',
	{big.NewInt(49), big.NewInt(183)}:  '2',
	{big.NewInt(49), big.NewInt(568)}:  '3',
	{big.NewInt(53), big.NewInt(277)}:  '4',
	{big.NewInt(53), big.NewInt(474)}:  '5',
	{big.NewInt(56), big.NewInt(332)}:  '6',
	{big.NewInt(56), big.NewInt(419)}:  '7',
	{big.NewInt(58), big.NewInt(139)}:  '8',
	{big.NewInt(58), big.NewInt(612)}:  '9',
	{big.NewInt(59), big.NewInt(365)}:  ':',
	{big.NewInt(59), big.NewInt(386)}:  ';',
	{big.NewInt(61), big.NewInt(129)}:  '<',
	{big.NewInt(61), big.NewInt(622)}:  '=',
	{big.NewInt(62), big.NewInt(372)}:  '>',
	{big.NewInt(62), big.NewInt(379)}:  '?',
	{big.NewInt(66), big.NewInt(199)}:  '@',
	{big.NewInt(66), big.NewInt(552)}:  'A',
	{big.NewInt(67), big.NewInt(84)}:   'B',
	{big.NewInt(67), big.NewInt(667)}:  'C',
	{big.NewInt(69), big.NewInt(241)}:  'D',
	{big.NewInt(69), big.NewInt(510)}:  'E',
	{big.NewInt(70), big.NewInt(195)}:  'F',
	{big.NewInt(70), big.NewInt(556)}:  'G',
	{big.NewInt(72), big.NewInt(254)}:  'H',
	{big.NewInt(72), big.NewInt(497)}:  'I',
	{big.NewInt(73), big.NewInt(72)}:   'J',
	{big.NewInt(73), big.NewInt(679)}:  'K',
	{big.NewInt(74), big.NewInt(170)}:  'L',
	{big.NewInt(74), big.NewInt(581)}:  'M',
	{big.NewInt(75), big.NewInt(318)}:  'N',
	{big.NewInt(75), big.NewInt(433)}:  'O',
	{big.NewInt(78), big.NewInt(271)}:  'P',
	{big.NewInt(78), big.NewInt(480)}:  'Q',
	{big.NewInt(79), big.NewInt(111)}:  'R',
	{big.NewInt(79), big.NewInt(640)}:  'S',
	{big.NewInt(80), big.NewInt(318)}:  'T',
	{big.NewInt(80), big.NewInt(433)}:  'U',
	{big.NewInt(82), big.NewInt(270)}:  'V',
	{big.NewInt(82), big.NewInt(481)}:  'W',
	{big.NewInt(83), big.NewInt(373)}:  'X',
	{big.NewInt(83), big.NewInt(378)}:  'Y',
	{big.NewInt(85), big.NewInt(35)}:   'Z',
	{big.NewInt(85), big.NewInt(716)}:  '[',
	{big.NewInt(86), big.NewInt(25)}:   '\\',
	{big.NewInt(86), big.NewInt(726)}:  ']',
	{big.NewInt(90), big.NewInt(21)}:   '^',
	{big.NewInt(90), big.NewInt(730)}:  '_',
	{big.NewInt(93), big.NewInt(267)}:  '`',
	{big.NewInt(93), big.NewInt(484)}:  'a',
	{big.NewInt(98), big.NewInt(338)}:  'b',
	{big.NewInt(98), big.NewInt(413)}:  'c',
	{big.NewInt(99), big.NewInt(295)}:  'd',
	{big.NewInt(99), big.NewInt(456)}:  'e',
	{big.NewInt(100), big.NewInt(364)}: 'f',
	{big.NewInt(100), big.NewInt(387)}: 'g',
	{big.NewInt(102), big.NewInt(267)}: 'h',
	{big.NewInt(102), big.NewInt(484)}: 'i',
	{big.NewInt(105), big.NewInt(369)}: 'j',
	{big.NewInt(105), big.NewInt(382)}: 'k',
	{big.NewInt(106), big.NewInt(24)}:  'l',
	{big.NewInt(106), big.NewInt(727)}: 'm',
	{big.NewInt(108), big.NewInt(247)}: 'n',
	{big.NewInt(108), big.NewInt(504)}: 'o',
	{big.NewInt(109), big.NewInt(200)}: 'p',
	{big.NewInt(109), big.NewInt(551)}: 'q',
	{big.NewInt(110), big.NewInt(129)}: 'r',
	{big.NewInt(110), big.NewInt(622)}: 's',
	{big.NewInt(114), big.NewInt(144)}: 't',
	{big.NewInt(114), big.NewInt(607)}: 'u',
	{big.NewInt(115), big.NewInt(242)}: 'v',
	{big.NewInt(115), big.NewInt(509)}: 'w',
	{big.NewInt(116), big.NewInt(92)}:  'x',
	{big.NewInt(116), big.NewInt(659)}: 'y',
	{big.NewInt(120), big.NewInt(147)}: 'z',
	{big.NewInt(120), big.NewInt(604)}: '{',
	{big.NewInt(125), big.NewInt(292)}: '|',
	{big.NewInt(125), big.NewInt(459)}: '}',
	{big.NewInt(126), big.NewInt(33)}:  '~',
	{big.NewInt(189), big.NewInt(297)}: 'А',
	{big.NewInt(189), big.NewInt(454)}: 'Б',
	{big.NewInt(192), big.NewInt(32)}:  'В',
	{big.NewInt(192), big.NewInt(719)}: 'Г',
	{big.NewInt(194), big.NewInt(205)}: 'Д',
	{big.NewInt(194), big.NewInt(546)}: 'Е',
	{big.NewInt(197), big.NewInt(145)}: 'Ж',
	{big.NewInt(197), big.NewInt(606)}: 'З',
	{big.NewInt(198), big.NewInt(224)}: 'И',
	{big.NewInt(198), big.NewInt(527)}: 'Й',
	{big.NewInt(200), big.NewInt(30)}:  'К',
	{big.NewInt(200), big.NewInt(721)}: 'Л',
	{big.NewInt(203), big.NewInt(324)}: 'М',
	{big.NewInt(203), big.NewInt(427)}: 'Н',
	{big.NewInt(205), big.NewInt(372)}: 'О',
	{big.NewInt(205), big.NewInt(379)}: 'П',
	{big.NewInt(206), big.NewInt(106)}: 'Р',
	{big.NewInt(206), big.NewInt(645)}: 'С',
	{big.NewInt(209), big.NewInt(82)}:  'Т',
	{big.NewInt(209), big.NewInt(669)}: 'У',
	{big.NewInt(210), big.NewInt(31)}:  'Ф',
	{big.NewInt(210), big.NewInt(720)}: 'Х',
	{big.NewInt(215), big.NewInt(247)}: 'Ц',
	{big.NewInt(215), big.NewInt(504)}: 'Ч',
	{big.NewInt(218), big.NewInt(150)}: 'Ш',
	{big.NewInt(218), big.NewInt(601)}: 'Щ',
	{big.NewInt(221), big.NewInt(138)}: 'Ъ',
	{big.NewInt(221), big.NewInt(613)}: 'Ы',
	{big.NewInt(226), big.NewInt(9)}:   'Ь',
	{big.NewInt(226), big.NewInt(742)}: 'Э',
	{big.NewInt(227), big.NewInt(299)}: 'Ю',
	{big.NewInt(227), big.NewInt(452)}: 'Я',
	{big.NewInt(228), big.NewInt(271)}: 'а',
	{big.NewInt(228), big.NewInt(480)}: 'б',
	{big.NewInt(229), big.NewInt(151)}: 'в',
	{big.NewInt(229), big.NewInt(600)}: 'г',
	{big.NewInt(234), big.NewInt(164)}: 'д',
	{big.NewInt(234), big.NewInt(587)}: 'е',
	{big.NewInt(235), big.NewInt(19)}:  'ж',
	{big.NewInt(235), big.NewInt(732)}: 'з',
	{big.NewInt(236), big.NewInt(39)}:  'и',
	{big.NewInt(236), big.NewInt(712)}: 'й',
	{big.NewInt(237), big.NewInt(297)}: 'к',
	{big.NewInt(237), big.NewInt(454)}: 'л',
	{big.NewInt(238), big.NewInt(175)}: 'м',
	{big.NewInt(238), big.NewInt(576)}: 'н',
	{big.NewInt(240), big.NewInt(309)}: 'о',
	{big.NewInt(240), big.NewInt(442)}: 'п',
	{big.NewInt(243), big.NewInt(87)}:  'р',
	{big.NewInt(243), big.NewInt(664)}: 'с',
	{big.NewInt(247), big.NewInt(266)}: 'т',
	{big.NewInt(247), big.NewInt(485)}: 'у',
	{big.NewInt(249), big.NewInt(183)}: 'ф',
	{big.NewInt(249), big.NewInt(568)}: 'х',
	{big.NewInt(250), big.NewInt(14)}:  'ц',
	{big.NewInt(250), big.NewInt(737)}: 'ч',
	{big.NewInt(251), big.NewInt(245)}: 'ш',
	{big.NewInt(251), big.NewInt(506)}: 'щ',
	{big.NewInt(253), big.NewInt(211)}: 'ъ',
	{big.NewInt(253), big.NewInt(540)}: 'ы',
	{big.NewInt(256), big.NewInt(121)}: 'ь',
	{big.NewInt(256), big.NewInt(630)}: 'э',
	{big.NewInt(257), big.NewInt(293)}: 'ю',
	{big.NewInt(257), big.NewInt(458)}: 'я',
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

func decryptPoint(C1, C2 Point, nb *big.Int) Point {
	nbC1 := scalarMult(nb, C1)                                   // nb * C1
	Pm := addPoints(C2, Point{nbC1.X, new(big.Int).Neg(nbC1.Y)}) // Pm = C2 - nb * C1
	return Pm
}

func DecryptWord(C1List, C2List []Point, nb *big.Int) []Point {
	PmList := []Point{}

	for i := range C1List {
		C1 := C1List[i]
		C2 := C2List[i]
		Pm := decryptPoint(C1, C2, nb)
		PmList = append(PmList, Pm)
	}
	return PmList
}

func PointsToCharacters(points []Point) string {
	var decryptedWord strings.Builder

	for _, decryptedPoint := range points {
		for point, letter := range alphabet {
			if decryptedPoint.X.Cmp(point.X) == 0 && decryptedPoint.Y.Cmp(point.Y) == 0 {
				decryptedWord.WriteRune(letter)

				break
			}
		}
	}

	return decryptedWord.String()
}
