package katas

import "strings"

func RomanLiteralToNumber(literal string) int {
	var score = 0
	var literals = strings.Split(literal, "")

	m := make(map[string]int)
	m["I"] = 1
	m["V"] = 5
	m["X"] = 10
	m["L"] = 50
	m["C"] = 100
	m["D"] = 500
	m["M"] = 1000

	for index := 0; index < len(literals); index++ {
		current := literals[index]
		score += m[current]

		if index > 0 && (m[literals[index-1]] < m[current]) {
			score -= 2 * m[literals[index-1]]
		}
	}

	return score
}
