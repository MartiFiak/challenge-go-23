package piscine

func Capitalize(s string) string {
	srune := []rune(s)
	srune2 := []rune{}
	for _, st := range srune {
		if st >= 65 && st <= 90 {
			srune2 = append(srune2, st+32)
			//ajout nouvelle case tableau
		} else {
			srune2 = append(srune2, st)

		}
	}
	str := ""
	for o, st := range srune2 {
		if o == 0 && st >= 97 && st <= 122 {
			str += string(st - 32)
			// str = str + string(i - 32)
		} else if st >= 97 && st <= 122 {
			if srune2[o-1] >= 48 && srune2[o-1] <= 57 {
				str += string(st)
			} else if srune2[o-1] >= 97 && srune2[o-1] <= 122 {
				str += string(st)
			} else {
				str += string(st - 32)
			}
		} else {
			str += string(st)
		}
	}
	return str
}
