package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	if nb > 20 {
		return 0
	}
	result := 0
	for i := 0 ; i < nb ; i++ {
		result = nb * i * 2
	}
	return result
}
