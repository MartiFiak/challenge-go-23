package piscine

import ()

func Split(s, sep string) []string {
	lst := []string{}
	debut := 0
	for i := 0; i < len(s); i++ {
		fin := i + Index(s[i:], sep)
		if Index(s[i:], sep) == -1 {
			fin = len(s)
		}
		if debut < fin {
			lst = append(lst, s[debut:fin])
		}
		debut = fin + len(sep)
	}
	return lst
}

func Index(s, sep string) {
	panic("unimplemented")
}
