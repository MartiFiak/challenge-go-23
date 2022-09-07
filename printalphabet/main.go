package main

import "github.com/01-edu/z01"

func main() {
	for lettre := 'z'; lettre != 'a'; lettre++ {
		z01.PrintRune(rune(lettre))
	}
	z01.PrintRune('\n')
}
