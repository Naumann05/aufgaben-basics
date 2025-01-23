package strings

// Erwartet einen string s und zählt, wie oft der Buchstabe 'A' in s vorkommt.
func CountA(s string) int {
	anzahl := 0
	for _, char := range s {
		if char == 'A' {
			anzahl++
		}
	}
	return anzahl
}

// Erwartet einen string s und einen Buchstaben c.
// Zählt, wie oft c in s vorkommt.
func CountChar(s string, c rune) int {

	result := 0
	for _, char := range s {
		if char == c {
			result++
		}
	}
	return result

}
