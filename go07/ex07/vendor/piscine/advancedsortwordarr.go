package piscine
	
func AdvancedSortWordArr(a []string, f func(a, b string) int) {
	for i := 0; i < elemLen(a) - 1; i++ {
		for j := i + 1; j < elemLen(a); j++ {
			if f(a[i], a[j]) == 1 {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

func Compare(a, b string) int {
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	return 0
}

func elemLen(strs []string) int {
	c := 0
	for range strs {
		c++
	}
	return c
}

/*
The Compare function can return three values (all of them type int):

0 is returned when the first string equals the second string (a==b).
+1 is returned when the first string is Lexicographically greater than the second string (a>b).
-1 is returned when the first string is Lexicographically smaller than the second string (a<b).
*/

