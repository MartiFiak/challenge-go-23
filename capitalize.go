package piscine

func Capitalize(s string) string {
	srune := []rune(s)
	srune2 := []rune{}
	for _, i := range srune {
		if i >= 65 && i <= 90 {
			srune2 = append(srune2, i+32)
		} else {
			srune2 = append(srune2, i)
		}
	}
	str := ""
	for v, i := range srune2 {
		if v == 0 && i >= 97 && i <= 122 {
			str += string(i - 32)
		} else if i >= 97 && i <= 122 {
			if srune2[v-1] >= 48 && srune2[v-1] <= 57 {
				str += string(i)
			} else if srune2[v-1] >= 97 && srune2[v-1] <= 122 {
				str += string(i)
			} else {
				str += string(i - 32)
			}
		} else {
			str += string(i)
		}
	}
	return str
}
