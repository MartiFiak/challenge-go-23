package piscine

func SplitWhiteSpaces(s string) []string {
	srune := []rune(s)
	lst := []string{}
	str := ""
	for i := 0; i < len(srune); i++ {
		if srune[i] == ' ' || srune[i] == '\t' || srune[i] == '\n' {
			if str != "" {
				lst = append(lst, str)
			}
			str = ""
		} else {
			str += string(srune[i])
		}
	}
	lst = append(lst, str)
	return lst
}
