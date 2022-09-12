package piscine

func IterativePower(nb int, power int) int {
	result := 1
	if power < 0 || nb > 70 {
		return 0
	}
	for i := power; i > 0; i-- {
		result *= nb
	}
	return result
}
