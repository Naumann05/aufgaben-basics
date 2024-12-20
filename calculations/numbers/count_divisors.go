package numbers

// Erwartet eine Zahl n >= 1 und liefert die Anzahl der Teiler dieser Zahl zurück.
func CountDivisors(n int) int {
	anzahl := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			anzahl++
		}
	}
	return anzahl
}

// muss mit X%Y gemaxht werde "%" ist der operator für Modulo Operationen diese geben den rest einer division aus.
