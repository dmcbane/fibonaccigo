package fibpkg

import (
	"math/big"
	"testing"
)

func TestGoldenRatioGenerator(t *testing.T) {
	gen := GoldenRatioGenerator()
	gen2 := GoldenRatioGenerator() // ensure the generators are independent
	expected := new(big.Rat)
	for i := uint32(1); i < uint32(200); i++ {
		expected.SetFrac(FibFastDoubling(i+1), FibFastDoubling(i))
		result := gen()
		gen2() // call the other generator to ensure no interaction
		if expected.Cmp(result) != 0 {
			t.Errorf("Expected gen [golden_ratio('%v')] to be %v, got %v", i, expected, result)
		}
	}
}

func BenchmarkGoldenRatio(b *testing.B) {
	gen := GoldenRatioGenerator()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gen()
	}
}
