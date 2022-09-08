package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	hello := "Hello World!"
	for _, s := range hello {
		z01.PrintRune(rune(s))
	}
}
