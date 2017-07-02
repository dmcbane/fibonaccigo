package fibpkg

import (
	"math/big"
)

// Fast doubling method. Faster than the matrix method.
// F(2n) = F(n) * (2*F(n+1) - F(n)).
// F(2n+1) = F(n+1)^2 + F(n)^2.
// This implementation is the non-recursive version. See
// https://www.nayuki.io/page/fast-fibonacci-algorithms
// for recursive and other implementations
func FibFastDoubling(n uint32) *big.Int {
	en := big.NewInt(int64(n))
	a := big.NewInt(0)
	b := big.NewInt(1)
	c := new(big.Int)
	d := new(big.Int)
	e := new(big.Int)
	f := new(big.Int)

	for i := uint32(31); i >= 0; i-- {
		c.Lsh(b, 1)
		c.Sub(c, a)
		c.Mul(c, a)
		d.Mul(a, a)
		e.Mul(b, b)
		b.Add(d, e)
		a = c
		// Advance by one if the bit is one
		if en.Bit(int(i)) != 0 {
			f.Add(a, b)
			a = b
			b = f
		}
	}
	return a
}
