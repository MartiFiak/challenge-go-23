package piscine

import "github.com/01-edu/z01"

func PrintComb() {
	for a := 48; a <= 57; a++ {
		for b := 48; b <= 57; b++ {
			for c := 48; c <= 57; c++ {
				if a != b && a != c && b != c && a < b && b < c {
					z01.PrintRune(rune(a))
					z01.PrintRune(rune(b))
					z01.PrintRune(rune(c))
					if a == 55 && b == 56 && c == 57 {
						z01.PrintRune('\n')
					} else {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
}
