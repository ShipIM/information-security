package utils

import (
	"math/big"
)

func Encode(N, e, C *big.Int) *big.Int {
	Ci := new(big.Int).Exp(C, e, N)
	res := new(big.Int)

	for Ci.Cmp(C) != 0 {
		res.Set(Ci)
		Ci = new(big.Int).Exp(Ci, e, N)
	}

	return res
}
