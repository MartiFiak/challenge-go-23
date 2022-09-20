package piscine

func SplitWhiteSpaces(s string) []string {
	n := " "
	wrd []string{}
	str := ""
	for a, i := range s {
		if i == " " || i == "\n" || i == "\t" || str != ""{
			wrd = append(wrd, str)
			str = ""
		} else {
			str += i
		}

	}
	return str
}