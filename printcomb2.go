package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	for a := 48; a <= 57; a++ {
		for b := 48; b <= 57; b++ {
			for c := 48; c <= 56; c++ {
				for d := 48; d <= 57; d++ {
					if a == c {
						if b < d {
							z01.PrintRune(a)
							z01.PrintRune(b)
							z01.PrintRune(32)
							z01.PrintRune(c)
							z01.PrintRune(d)
							if a == allDigits[9] && b == allDigits[8] && c == allDigits[9] && d == allDigits[9] {
								z01.PrintRune('\n')
							} else {
								z01.PrintRune(',')
								z01.PrintRune(32)
							}
						}
					} else if a < c {
						z01.PrintRune(a)
						z01.PrintRune(b)
						z01.PrintRune(32)
						z01.PrintRune(c)
						z01.PrintRune(d)
						if a == 57 && b == 56 && c == 57 && d == 57 {
							z01.PrintRune('\n')
						} else {
							z01.PrintRune(',')
							z01.PrintRune(32)
						}
					}
				}
			}
		}
	}
}
