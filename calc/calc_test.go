package calc

import "testing"

func TestAddition(t *testing.T) {
	var calculator Addition
	assertEqual(t, 0, calculator.Calculate(0, 0))
	assertEqual(t, 1, calculator.Calculate(0, 1))
	assertEqual(t, 2, calculator.Calculate(1, 1))
	assertEqual(t, 6, calculator.Calculate(3, 4))
	assertEqual(t, -5, calculator.Calculate(2, -7))
}

func assertEqual(t *testing.T, expected, actual any) {
	t.Helper()
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
