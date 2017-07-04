package fibpkg

import (
	"testing"
)

func TestFibonacciGenerator(t *testing.T) {
	gen := FibonacciGenerator()
	for i := uint32(0); i < uint32(200); i++ {
		expected := FibFastDoubling(i)
		result := gen()
		if expected.Cmp(result) != 0 {
			t.Errorf("Expected gen [fib('%v')] to be %v, got %v", i, expected, result)
		}
	}
}

func BenchmarkFibonacciGenerator(b *testing.B) {
	gen := FibonacciGenerator()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		gen()
	}
}
