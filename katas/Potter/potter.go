package Potter

import "sort"

type Combination struct {
	value []int
}
type ByValueLength []Combination

func (a ByValueLength) Len() int { return len(a) }
func (a ByValueLength) Less(i, j int) bool {
	return len(a[i].value) >= 3 && len(a[i].value) < len(a[j].value)
}
func (a ByValueLength) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func isInArray(value int, values []int) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

func CalculatePrice(books []int) float64 {
	var price = 0.0
	const costPerBook = 8
	discounts := map[int]float64{1: 1, 2: 0.95, 3: 0.9, 4: 0.8, 5: 0.75}
	combinations := []Combination{Combination{value: []int{}}}

	for _, book := range books {
		var added = false
		for j := 0; j < len(combinations); j++ {
			if !isInArray(book, combinations[j].value) {
				combinations[j].value = append(combinations[j].value, book)
				sort.Sort(ByValueLength(combinations))
				added = true
				break
			}
		}
		if !added {
			combinations = append(combinations, Combination{value: []int{book}})
		}
	}

	for _, combination := range combinations {
		price += float64(float64(len(combination.value)) * float64(costPerBook) * discounts[len(combination.value)])
	}
	return price
}
