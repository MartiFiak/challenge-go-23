package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args
	for i, c := range argument {
		if i != 0 {
			for _, k := range c {
				z01.PrintRune(k)
			}
			z01.PrintRune('\n')
		}
	}
}
