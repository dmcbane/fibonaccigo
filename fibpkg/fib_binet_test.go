package fibpkg

import "testing"

func TestFibBinet(t *testing.T) {
	param := uint32(0)
	expected := float64(0.0)
	result := FibBinet(param)
	if expected != result {
		t.Errorf("Expected FibBinet('%v') to be %v, got %v", param, expected, result)
	}
	param = 1
	expected = 1.0
	result = FibBinet(param)
	if expected != result {
		t.Errorf("Expected FibBinet('%v') to be %v, got %v", param, expected, result)
	}
	param = 2
	expected = 1.0
	result = FibBinet(param)
	if expected != result {
		t.Errorf("Expected FibBinet('%v') to be %v, got %v", param, expected, result)
	}
	param = 75
	expected = 2111485077978050.0
	result = FibBinet(param)
	if expected != result {
		t.Errorf("Expected FibBinet('%v') to be %v, got %v", param, expected, result)
	}
}
func TestFibBinetBroken(t *testing.T) {
	var param uint32
	var expected float64
	var result float64
	param = 76
	expected = 3416454622906707.0
	result = FibBinet(param)
	if expected == result {
		t.Errorf("Expected FibBinet('%v') to be %v, got %v", param, expected, result)
	}
}

func BenchmarkFibBinet75(b *testing.B) {
	// b.ResetTimer()
	for i := 0; i < b.N; i++ {
		FibBinet(75)
	}
}
