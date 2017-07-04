package fibpkg

import (
	"math/big"
)

func GoldenRatioGenerator() func() *big.Rat {
	f0 := big.NewInt(1)
	f1 := big.NewInt(1)
	f2 := big.NewInt(2)
	return func() *big.Rat {
		tmp := new(big.Rat)
		tmp.SetFrac(f1, f0)
		f0.Set(f1)
		f1.Set(f2)
		f2.Add(f0, f1)
		return tmp
	}
}
