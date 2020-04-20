package katas

import (
	"math"
)

func isEvenNumber(number int) bool {
	return math.Mod(math.Abs(float64(number)), 2) == float64(0)
}

func FindOutlier(numbers []int) int {
	odd := make([]int, 0)
	even := make([]int, 0)

	for index := 0; index < len(numbers); index++ {
		if isEvenNumber(numbers[index]) {
			even = append(even, numbers[index])
		} else {
			odd = append(odd, numbers[index])
		}

		if len(even) == 1 && len(odd) > 1 {
			return even[0]
		} else if len(odd) == 1 && len(even) > 1 {
			return odd[0]
		}
	}

	return 0
}
