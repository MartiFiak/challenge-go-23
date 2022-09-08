package piscine

func StrRev(s string) string {
	sString := []rune(s)
	f := ""
	for j := len(sString) - 1; j >= 0; j-- {
		f += string(sString[j])
	}
	return f
}
