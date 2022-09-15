package piscine

func BasicJoin(elems []string) string {
	var s string
	for i := range elems {
		s = s + elems[i]
	}
	return s
}
