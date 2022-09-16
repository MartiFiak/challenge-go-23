package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argt := os.Args
	for i := len(argt) - 1; i > 0; i-- {
		if i != 0 {
			for _, k := range argt[i] {
				z01.PrintRune(k)
			}
			z01.PrintRune('\n')
		}
	}
}
