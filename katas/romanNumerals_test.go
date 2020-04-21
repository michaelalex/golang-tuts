package katas

import "testing"

func TestEmptyRomanLiteralToNumber(t *testing.T) {
	var expected = 0
	var value = RomanLiteralToNumber("")

	if value != expected {
		t.Errorf("Error expected %d, actual: %d", expected, value)
	}
}

func TestSingleRomanLiteralToNumber(t *testing.T) {
	var expected = 1
	var value = RomanLiteralToNumber("I")

	if value != expected {
		t.Errorf("Error expected %d, actual: %d", expected, value)
	}
}

func TestFiveRomanLiteralToNumber(t *testing.T) {
	var expected = 5
	var value = RomanLiteralToNumber("V")

	if value != expected {
		t.Errorf("Error expected %d, actual: %d", expected, value)
	}
}

func TestTenRomanLiteralToNumber(t *testing.T) {
	var expected = 5
	var value = RomanLiteralToNumber("V")

	if value != expected {
		t.Errorf("Error expected %d, actual: %d", expected, value)
	}
}

func TestNineRomanLiteralToNumber(t *testing.T) {
	var expected = 9
	var value = RomanLiteralToNumber("IX")

	if value != expected {
		t.Errorf("Error expected %d, actual: %d", expected, value)
	}
}

func TestEightRomanLiteralToNumber(t *testing.T) {
	var expected = 8
	var value = RomanLiteralToNumber("VIII")

	if value != expected {
		t.Errorf("Error expected %d, actual: %d", expected, value)
	}
}

func TestThirtyFourRomanLiteralToNumber(t *testing.T) {
	var expected = 34
	var value = RomanLiteralToNumber("XXXIV")

	if value != expected {
		t.Errorf("Error expected %d, actual: %d", expected, value)
	}
}

func TestNintyRomanLiteralToNumber(t *testing.T) {
	var expected = 900
	var value = RomanLiteralToNumber("CM")

	if value != expected {
		t.Errorf("Error expected %d, actual: %d", expected, value)
	}
}
