package test

import "math"

var total int

func BonusCoklat(val float64) int {
	if total == 0 {
		total += int(val)
	}

	bonus := int(math.Floor(val / 5))
	total += bonus
	
	if bonus >= 5 {
		BonusCoklat(float64(bonus))
	}

	return total
}
