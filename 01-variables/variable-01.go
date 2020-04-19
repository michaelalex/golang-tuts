package variables

func GetTruthy() bool {
	return true
}

func GetCoupleOfNumbers() (int, int) {
	var b, c int = 1, 2
	return b, c
}

func GetHelloWorld() string {
	value := "Hello World"
	return value
}
