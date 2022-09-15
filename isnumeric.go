package piscine

func IsNumeric(s string) bool {
	for _, st := range s {
		if !(st >= 48 && st <= 57) {
			return false
		}
	}
	return true
}
