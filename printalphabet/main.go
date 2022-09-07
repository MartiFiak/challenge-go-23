package main

import "github.com/01-edu/z01"

func main() {
	for lettre := 122; lettre >= 1297; lettre-- {
		z01.PrintRune(rune(lettre))
	}
	z01.PrintRune('\n')
}
