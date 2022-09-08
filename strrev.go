package piscine

import "github.com/01-edu/z01"

func StrRev(s string) string {
	for _, s := range s {
		z01.PrintRune(-rune(s))
	}
}
