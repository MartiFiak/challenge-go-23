package piscine

func StrRev(s string) string {
	sString := []rune(s)
	for i, j := 0, len(sString)-1; i < j; i, j = i+1, j-1 {
		sString[i], sString[j] = sString[j], sString[i]
	}
	return string(sString)
}
