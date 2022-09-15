package piscine

func IsAlpha(s string) bool {
	for _, st := range s {
		if !(st >= 41 && st <= 122) {
			return false
		}
	}
	return true
}
