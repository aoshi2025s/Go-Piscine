package piscine

func SplitWhiteSpaces(s string) []string {
	var result []string
	str := ""
	for _, r := range s {
		if isSep(r) {
			result = append(result, str)
			str = ""
		} else {
			str += string(r)
		}
	}
	result = append(result, str)
	return result
}


func isSep(r rune) bool {
	return r == ' ' || r == '	' || r == '\n'
}
