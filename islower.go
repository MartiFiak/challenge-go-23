package piscine

func IsLower(s string) bool {
	for _, st := range s {
		if !(st >= 97 && st <= 122) {
			return false
		}
	}
	return true
}
