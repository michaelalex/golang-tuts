package variables

func GetTruthy() bool {
	return true
}

func GetCoupleOfNumbers() (int, int) {
	var b, c int = 1, 2
	return b, c
}

const value string = "Hello World"

func GetHelloWorld() string {
	return value
}
