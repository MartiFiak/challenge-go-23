package piscine

func IterativePower(nb int, power int) int {
	if power < 0 || power > 20{
		return 0
	}
	if nb < 0 || nb > 20{
		return 0
	}
	result := 1
	for i := 1; i < nb; i++ {
		result *= nb
	}
	return result
}
