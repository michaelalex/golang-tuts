package katas

import "strings"

type numeral struct {
	letter string
	value  int
}

type replacementNumeral struct {
	value       string
	replacement string
}

func CalculateRomanNumeral(romanNumerals string) int {
	var numerals = strings.Split(romanNumerals, "")

	var replacements = []replacementNumeral{
		{value: "VV", replacement: "X"},
		{value: "LL", replacement: "C"},
		{value: "DD", replacement: "M"},
		{value: "IIII", replacement: "IV"},
		{value: "XXXX", replacement: "XL"},
		{value: "CCCC", replacement: "CD"},
	}

	for index := 0; index < len(numerals); index++ {
		var numeral = findValue(numerals[index], numeralValues)

	}

	return 0
}
