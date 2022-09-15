package piscine

func IsPrintable(s string) bool {
	for _, st := range s {
		if !(st >= 32 && st <= 126) {
			return false
		}
	}
	return true
}
