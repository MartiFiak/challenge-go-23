package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for a := 48; a <= 57; a++ {
		for b := 48; b <= 57; b++ {
			for c := 48; c <= 56; c++ {
				for d := 49; d <= 57; d++ {
					if a <= c && b < d {
						z01.PrintRune(rune(a))
						z01.PrintRune(rune(b))
						z01.PrintRune(' ')
						z01.PrintRune(rune(c))
						z01.PrintRune(rune(d))
						if a == 57 && b == 56 && c == 57 && d == 57 {
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
}