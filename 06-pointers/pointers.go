package pointers

func UpdateValueWithoutPointer(value int, updateTo int) {
	value = updateTo
}

func UpdateValue(value *int, updateTo int) {
	*value = updateTo
}
