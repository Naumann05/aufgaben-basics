package strings

// Erwartet einen String s und einen Buchstaben c.
// Prüft, ob c in s vorkommt.
func Contains(s string, c byte) bool {
	for char := range s {
		if s[char] == c {
			return true
		}
	}
	return false
}
