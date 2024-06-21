package piscine

func CountArg(params []string) int {
	c := 0
	for range params {
		c++
	}
	return c
}
