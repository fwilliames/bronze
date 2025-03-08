package utils

import (
	"math"
	"testing"
)

func TestSoma(t *testing.T) {
	values := []float64{10.5, 20.0, 15.75, 30.25, 5.99}

	resultado := Sum(values)
	esperado := 82.49

	if resultado != esperado {
		t.Errorf("Soma(2, 3) = %f; esperado %f", resultado, esperado)
	}
}
func TestSum_EmptySlice(t *testing.T) {
	values := []float64{} // Lista vazia
	expected := 0.0

	result := Sum(values)

	if result != expected {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}

func TestSum_WithNaN(t *testing.T) {
	values := []float64{1.0, math.NaN(), 2.0}

	result := Sum(values)

	if !math.IsNaN(result) {
		t.Errorf("Expected NaN, but got %v", result)
	}
}

func TestSum_WithInfinity(t *testing.T) {
	values := []float64{1.0, math.Inf(1), 2.0}

	result := Sum(values)

	if !math.IsInf(result, 1) {
		t.Errorf("Expected +Inf, but got %v", result)
	}
}
func TestMultiplica(t *testing.T) {
	resultado := Prod(2, 3)
	esperado := 6.00

	if resultado != esperado {
		t.Errorf("Multiplica(2, 3) = %f; esperado %f", resultado, esperado)
	}
}
