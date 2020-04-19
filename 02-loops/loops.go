package loops

func BasicIterate(iterations int, callback func(currentIndex int)) {
	i := 1
	for i <= iterations {
		callback(i)
		i++
	}
}

func Iterate(iterations int, callback func(currentIndex int)) {
	for index := 1; index <= iterations; index++ {
		callback(index)
	}
}

func InfiniteLoop(callback func()) {
	for {
		callback()
	}
}
