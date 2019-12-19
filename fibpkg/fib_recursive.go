package fibpkg

import (
	"math/big"
)

func fibacc(n, prev, acc *big.Int) *big.Int {
	if n.Cmp(big.NewInt(0)) <= 0 {
		return acc
	} else {
		return fibacc(n.Sub(n, big.NewInt(1)), acc, prev.Add(prev, acc))
	}
}

func FibRecursive(n *big.Int) *big.Int {
	return fibacc(big.NewInt(n.Int64()), big.NewInt(1), big.NewInt(0))
}
