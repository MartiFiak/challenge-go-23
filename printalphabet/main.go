package main

import "github.com/01-edu/z01"

func main() {
	for lettre := 97; lettre <= 122; lettre++ {
		z01.PrintRune(rune(lettre))
	}
}
