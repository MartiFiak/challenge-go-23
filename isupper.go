package piscine

func IsUpper(s string) bool {
	for _, st := range s {
		if !(st >= 65 && st <= 90) {
			return false
		}
	}
	return true
}
