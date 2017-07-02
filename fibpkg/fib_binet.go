package fibpkg

import (
	"math"
)

// Calculate fibonacci numbers with internal representations
// this version looses accuracy after the 75th fibonacci
// number due to limits of f64
func FibBinet(n uint32) float64 {
	sqrt5 := math.Sqrt(5.0)
	return math.Floor(math.Pow(math.Phi, float64(n))/sqrt5 + 0.5)
}
