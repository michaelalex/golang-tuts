package arrays

func IterateRange(arr []int, callback func(currentValue int)) {
	for _, num := range arr {
		callback(num)
	}
}

func IterateRangeWithIndex(arr []int, callback func(currentValue int, currentIndex int)) {
	for i, num := range arr {
		callback(num, i)
	}
}
