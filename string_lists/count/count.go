package count

// Erwartet eine Liste von Strings sowie einen String, der gezählt werden soll.
// Liefer die Anzahl der Vorkommen des gesuchten Strings in der Liste.
func Count(strings []string, search string) int {
	anzahl := 0
	for _, s := range strings {
		if s == search {
			anzahl++
		}
	}
	return anzahl
}
