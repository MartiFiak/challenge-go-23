package piscine

func AlphaCount(s string) int {
	count := 0
	for _, st := range s {
		if (st >= 65 && st <= 90) || (st >= 97 && st <= 122) {
			count = count + 1
		}
	}
	return count
}
