package arrays

func makeArray(capacity int) []int {
	return make([]int, capacity)
}

func makeTwoDimArray(xCapacity int, yCapacity int) [][]int {
	return make([][]int, xCapacity, yCapacity)
}
