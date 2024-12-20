package lists

import "math"

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.
// ZUSATZBEDINGUNG: Diese Funktion darf keine Schleife verwenden.
func MinListRecursive(nums []int) int {
	// TODO
	smallest := math.MaxInt
	if len(nums) == 0 {
		return 0
	}
	for _, num := range nums {
		if num < smallest {
			smallest = num
		}
	}
	return smallest
}

// REMARKS
// - Diese Funktion nennt man "rekursiv".
// - Rekursion ist ein größeres Thema, das in der Vorlesung separat behandelt wird.
//   Um die Denkweise frühzeitig zu üben, gibt es aber gelegentlich auch vorab Aufgaben dazu.
