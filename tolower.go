package piscine

func ToLowerer(s string) string {
	str := ""
	srune := []rune(s)
	for _, st := range srune {
		if st >= 65 && st <= 90 {
			str += string(st + 32)
		} else {
			str += string(st)
		}
	}
	return str
}
