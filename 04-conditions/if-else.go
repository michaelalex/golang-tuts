package conditions

func OddEvenCallback(value int, oddCallback func(value int), evenCallback func(value int)) {
	if value%2 == 0 {
		evenCallback(value)
	} else {
		oddCallback(value)
	}
}
