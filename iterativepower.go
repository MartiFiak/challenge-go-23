package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	if nb > 70 {
		return 0
	}
	result := 1
	for i := power; i > 0; i-- {
		result *= nb
	}
	return result
}
