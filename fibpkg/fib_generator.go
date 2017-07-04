package fibpkg

import (
	"math/big"
)

func FibonacciGenerator() func() *big.Int {
	f0 := new(big.Int)
	f1 := big.NewInt(1)
	f2 := big.NewInt(1)
	return func() *big.Int {
		tmp := new(big.Int)
		tmp.Set(f0)
		f0.Set(f1)
		f1.Set(f2)
		f2.Add(f0, f1)
		return tmp
	}
}
