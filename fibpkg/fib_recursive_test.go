package fibpkg

import (
	"math/big"
	"testing"
)

func TestFibRecursive(t *testing.T) {
	param := big.NewInt(0)
	expected := big.NewInt(0)
	result := FibRecursive(param)
	if expected.Cmp(result) != 0 {
		t.Errorf("Expected FibRecursive('%v') to be %v, got %v", param, expected, result)
	}
	param = big.NewInt(1)
	expected = big.NewInt(1)
	result = FibRecursive(param)
	if expected.Cmp(result) != 0 {
		t.Errorf("Expected FibRecursive('%v') to be %v, got %v", param, expected, result)
	}
	param = big.NewInt(2)
	expected = big.NewInt(1)
	result = FibRecursive(param)
	if expected.Cmp(result) != 0 {
		t.Errorf("Expected FibRecursive('%v') to be %v, got %v", param, expected, result)
	}
	param = big.NewInt(75)
	expected = big.NewInt(2111485077978050)
	result = FibRecursive(param)
	if expected.Cmp(result) != 0 {
		t.Errorf("Expected FibRecursive('%v') to be %v, got %v", param, expected, result)
	}
}

func BenchmarkFibRecursive75(b *testing.B) {
	param := big.NewInt(75)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FibRecursive(param)
	}
}
