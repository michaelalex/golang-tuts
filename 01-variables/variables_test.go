package variables

import "testing"

func TestGetTruthy(t *testing.T) {
	var truthy = GetTruthy()
	if truthy != true {
		t.Error("Truthy expected to be true")
	}
}

// func GetCoupleOfNumbers() (int, int) {
// 	var b, c int = 1, 2
// 	return b, c
// }

// const value string = "Hello World"

// func GetHelloWorld() string {
// 	return value
// }
