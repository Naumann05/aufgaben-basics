package lists

import (
	"math"
)

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.
func MinList(nums []int) int {
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
