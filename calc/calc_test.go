package calc

import "testing"

func TestAddition(t *testing.T) {
	assertEqual(t, 0, Addition{}.Calculate(0, 0))
	assertEqual(t, 1, Addition{}.Calculate(0, 1))
	assertEqual(t, 2, Addition{}.Calculate(1, 1))
	assertEqual(t, 7, Addition{}.Calculate(3, 4))
	assertEqual(t, -5, Addition{}.Calculate(2, -7))
}

func assertEqual(t *testing.T, expected, actual any) {
	if actual != expected {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
