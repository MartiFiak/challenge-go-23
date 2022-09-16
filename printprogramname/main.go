package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arguments := []rune(os.Args[0])
	for _, i := range arguments[2:] {
		z01.PrintRune(i)
	}
	z01.PrintRune('\n')
}
