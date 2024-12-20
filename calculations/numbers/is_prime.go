package numbers

// Erwartet eine Zahl n und prüft, ob n eine Primzahl ist.
func IsPrime(n int) bool {
	anzahl := 0
	var Prime bool
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			anzahl++
		}
	}
	if anzahl == 2 {
		Prime = true
	}
	return Prime
}
