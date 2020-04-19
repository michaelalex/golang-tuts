package functions

func AppendString(initialString string) func(string) string {
	t := initialString
	c := func(b string) string {
		t = t + " " + b
		return t
	}
	return c
}
