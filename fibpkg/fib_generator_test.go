package fibpkg

import (
	"testing"
)

func TestFibonacciGenerator(t *testing.T) {
	gen := FibonacciGenerator()
	gen2 := FibonacciGenerator() // ensure the generators are independent
	for i := uint32(0); i < uint32(200); i++ {
		expected := FibFastDoubling(i)
		result := gen()
		gen2() // call the other generator to ensure no interaction
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
