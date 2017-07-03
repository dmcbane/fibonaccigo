package fibpkg

import (
	"math/big"
)

// Fast doubling method. Faster than the matrix method.
// F(2n) = F(n) * (2*F(n+1) - F(n)).
// F(2n+1) = F(n+1)^2 + F(n)^2.
// This implementation is the non-recursive version. See
// https://www.nayuki.io/page/fast-fibonacci-algorithms
func FibFastDoubling(n uint32) *big.Int {
	en := big.NewInt(int64(n))
	two := big.NewInt(2)
	a := big.NewInt(0)
	b := big.NewInt(1)
	c := new(big.Int)
	d := new(big.Int)
	e := new(big.Int)
	f := new(big.Int)
	g := new(big.Int)
	h := new(big.Int)

	for i := 31; i >= 0; i-- {
		// d = a * (b * 2 - a);
		//         |  f  |
		//         |    g    |
		// f = b * 2
		f.Mul(b, two)
		// g = f - a
		g.Sub(f, a)
		// d = a * g
		d.Mul(a, g)
		// e = a * a + b * b;
		//     | g |   | h |
		//
		// g = a * a
		g.Mul(a, a)
		// h = b * b
		h.Mul(b, b)
		// e = g + h
		e.Add(g, h)
		// a = d
		a.Set(d)
		// b = e
		b.Set(e)
		// Advance by one if the bit is one
		// if (((n >> i) & 1) != 0) {
		if en.Bit(i) != 0 {
			// c = a + b;
			c.Add(a, b)
			// a = b
			a.Set(b)
			// b = c
			b.Set(c)
		}
	}
	return a
}
