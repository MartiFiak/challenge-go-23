package piscine

func ToUpper(s string) string {
	str := ""
	srune := []rune(s)
	for _, st := range srune {
		if st >= 97 && st <= 122 {
			str += string(st - 32)
		} else {
			str += string(st)
		}
	}
	return str
}
