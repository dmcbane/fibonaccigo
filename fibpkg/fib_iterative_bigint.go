package fibpkg

import (
	"math/big"
)

// Calculate large fibonacci numbers naive iterative
// algorithm using BigUint
func FibIterativeBigint(n uint32) *big.Int {
	f0 := big.NewInt(0)
	f1 := big.NewInt(1)
	f2 := big.NewInt(0)
	for i := uint32(0); i < n; i++ {
		f2.Add(f0, f1)
		f0 = f1
		f1 = f2
	}
	return f0
}
