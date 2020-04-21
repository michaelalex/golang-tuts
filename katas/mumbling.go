package katas

import "strings"

func Mumble(value string) string {
	var mumble = make([]string, 0)
	characters := strings.Split(value, "")

	for index := 0; index < len(characters); index++ {
		mumble = append(mumble, strings.ToUpper(characters[index])+strings.Repeat(strings.ToLower(characters[index]), index))
	}

	return strings.Join(mumble, "-")
}
