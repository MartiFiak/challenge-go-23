package piscine

func LastRune(s string) rune {
	st := []rune(s)
	return st[len(st)-1]
}
