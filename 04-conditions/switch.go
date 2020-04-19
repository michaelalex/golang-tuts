package conditions

func GetNumberAsString(value int) string {
	switch value {
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	default:
		return "unknown value"
	}
}

func WhatAmI(intFace interface{}) string {
	switch intFace.(type) {
	case bool:
		return "I'm a bool"
	case int:
		return "I'm an int"
	case string:
		return "I'm a string"
	default:
		return "Don't know type"
	}
}
