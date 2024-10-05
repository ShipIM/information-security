package utils

import "math/big"

var (
	zero = big.NewInt(0)
	one  = big.NewInt(1)
)

func ExtendedGCD(a, b *big.Int) (*big.Int, *big.Int, *big.Int) {
	if a.Cmp(zero) == 0 {
		return new(big.Int).Set(b), zero, one
	}

	gcd, x, y := ExtendedGCD(new(big.Int).Mod(b, a), a)

	return gcd, new(big.Int).Sub(y, new(big.Int).Mul(new(big.Int).Div(b, a), x)), new(big.Int).Set(x)
}
