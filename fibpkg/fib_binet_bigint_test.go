package fibpkg

import (
	"math/big"
	"testing"
)

func TestFibBinetBigint(t *testing.T) {
	param := uint32(0)
	expected := big.NewInt(0)
	result := FibBinetBigint(param)
	if expected.Cmp(result) != 0 {
		t.Errorf("Expected FibBinetBigint('%v') to be %v, got %v", param, expected, result)
	}
	param = 1
	expected = big.NewInt(1)
	result = FibBinetBigint(param)
	if expected.Cmp(result) != 0 {
		t.Errorf("Expected FibBinetBigint('%v') to be %v, got %v", param, expected, result)
	}
	param = 2
	expected = big.NewInt(1)
	result = FibBinetBigint(param)
	if expected.Cmp(result) != 0 {
		t.Errorf("Expected FibBinetBigint('%v') to be %v, got %v", param, expected, result)
	}
	param = 75
	expected.SetString("2111485077978050", 10)
	result = FibBinetBigint(param)
	if expected.Cmp(result) != 0 {
		t.Errorf("Expected FibBinetBigint('%v') to be %v, got %v", param, expected, result)
	}
	param = 76
	expected.SetString("3416454622906707", 10)
	result = FibBinetBigint(param)
	if expected.Cmp(result) != 0 {
		t.Errorf("Expected FibBinetBigint('%v') to be %v, got %v", param, expected, result)
	}
	param = 281
	expected.SetString("23770696554372451866815101694984845480039225387896643963981", 10)
	result = FibBinetBigint(param)
	if expected.Cmp(result) != 0 {
		t.Errorf("Expected FibBinetBigint('%v') to be %v, got %v", param, expected, result)
	}
}

func TestFibBinetBigintBroken(t *testing.T) {
	var param uint32
	expected := new(big.Int)
	result := new(big.Int)
	param = 282
	expected.SetString("38461794961234640015759308940939757590587318989278841661816", 10)
	result = FibBinetBigint(param)
	if expected == result {
		t.Errorf("Expected FibBinet('%v') to be %v, got %v", param, expected, result)
	}
}

func BenchmarkFibBinetBigint75(b *testing.B) {
	// b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FibBinetBigint(75)
	}
}

func BenchmarkFibBinetBigint281(b *testing.B) {
	// b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FibBinetBigint(281)
	}
}
