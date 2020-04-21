package katas

import "testing"

func TestMumbleWithA(t *testing.T) {
	var expected = "A"
	var result = Mumble("a")
	if result != expected {
		t.Errorf("Failed, expected %s ... actual %s", expected, result)
	}
}

func TestMumbleWithABC(t *testing.T) {
	var expected = "A-Bb-Ccc"
	var result = Mumble("abC")
	if result != expected {
		t.Errorf("Failed, expected %s ... actual %s", expected, result)
	}
}

func TestMumbleWithaBcD(t *testing.T) {
	var expected = "A-Bb-Ccc-Dddd"
	var result = Mumble("aBCd")
	if result != expected {
		t.Errorf("Failed, expected %s ... actual %s", expected, result)
	}
}

func TestMumbleWithQwerty(t *testing.T) {
	var expected = "Q-Ww-Eee-Rrrr-Ttttt-Yyyyyy"
	var result = Mumble("QWERTY")
	if result != expected {
		t.Errorf("Failed, expected %s ... actual %s", expected, result)
	}
}
