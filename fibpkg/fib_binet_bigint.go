package fibpkg

import (
	"math"
	"math/big"
)

func approx_sqrt(value *big.Float) *big.Float {
	// We'll do computations with 200 bits of precision in the mantissa.
	const prec = 200
	// Since Newton's Method doubles the number of correct digits at each
	// iteration, we need at least log_2(prec) steps.
	steps := int(math.Log2(prec))

	// Compute the square root of value using Newton's Method. We start with
	// an initial estimate for sqrt(2), and then iterate:
	//     x_{n+1} = 1/2 * ( x_n + (2.0 / x_n) )

	// Initialize values we need for the computation.
	half := new(big.Float).SetPrec(prec).SetFloat64(0.5)

	// Use 1 as the initial estimate.
	x := new(big.Float).SetPrec(prec).SetInt64(1)

	// We use t as a temporary variable. There's no need to set its precision
	// since big.Float values with unset (== 0) precision automatically assume
	// the largest precision of the arguments when used as the result (receiver)
	// of a big.Float operation.
	t := new(big.Float)

	// Iterate.
	for i := 0; i <= steps; i++ {
		t.Quo(value, x) // t = value / x_n
		t.Add(x, t)     // t = x_n + (value / x_n)
		x.Mul(half, t)  // x_{n+1} = 0.5 * t
	}
	return x
}

func powBigFloat(a *big.Float, n uint32) *big.Float {
	tmp := new(big.Float)
	tmp.Copy(a)
	res := big.NewFloat(1.0)
	for n > 0 {
		temp := new(big.Float)
		if n%2 == 1 {
			temp.Mul(res, tmp)
			res = temp
		}
		temp = new(big.Float)
		temp.Mul(tmp, tmp)
		tmp = temp
		n /= 2
	}
	return res
}

// Calculate fibonacci numbers with big representations
func FibBinetBigint(n uint32) *big.Int {
	sqrt5 := approx_sqrt(big.NewFloat(5.0))
	phi := new(big.Float)
	phi.Add(big.NewFloat(1.0), sqrt5)
	phi.Quo(phi, big.NewFloat(2.0))
	result := powBigFloat(phi, n)
	result.Quo(result, sqrt5)
	result.Add(result, big.NewFloat(0.5))
	rounded := new(big.Int)
	result.Int(rounded)
	return rounded
}
