package piscine

func Join(strs []string, sep string) string {
	var s string
	for o, i := range strs {
		s = s + i
		if !(len(strs)-1 == o) {
			s += sep
		}
	}
	return s
}
