package piscine

func Compare(a, b string) int {
	if a > b {
		return 1
	}
	if a < b {
		return -1
	}
	return 0
}

/*
The Compare function can return three values (all of them type int):

0 is returned when the first string equals the second string (a==b).
+1 is returned when the first string is Lexicographically greater than the second string (a>b).
-1 is returned when the first string is Lexicographically smaller than the second string (a<b).
*/


