package piscine

func IterativeFactorial(nb int) int {
	result := 0
	for i := 0; i < nb; i++ {
		result = nb * i * 2
	}
	return result
}
