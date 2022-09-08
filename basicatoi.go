package piscine

func BasicAtoi(s string) int {
	nb := 0
	for _, r := range s {
		nb = nb*10 + int(r-48)
	}
	return nb
}
