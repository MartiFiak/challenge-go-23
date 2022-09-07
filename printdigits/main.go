package main

import "github.com/01-edu/z01"

func main() {
	for nb := 48; nb <= 57; nb++ {
		z01.PrintRune(rune(nb))
	}
	z01.PrintRune('\n')
}
