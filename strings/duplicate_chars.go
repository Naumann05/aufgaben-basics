package strings

// Erwartet einen String s und liefert einen neuen String,
// in dem jeder Buchstabe aus s zweimal hintereinander vorkommt.
func DuplicateChars(s string) string {
	out := ""
	for _, char := range s {
		out += string(char)
		out += string(char)
	}
	return out
}
