package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 {
		return 0
	}
	if nb > 20 {
		return 0
	}
	result := 0
	for i := 0 ; i < nb + 3 ; i++ {
		result = nb * i
	}
	return result
}
