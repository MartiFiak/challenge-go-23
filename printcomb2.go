package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	allDigits := [10]rune{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}
	for _, a := range allDigits {
		for _, b := range allDigits {
			for _, c := range allDigits {
				for _, d := range allDigits {
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
						if a == allDigits[9] && b == allDigits[8] && c == allDigits[9] && d == allDigits[9] {
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
