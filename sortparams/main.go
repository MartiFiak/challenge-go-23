package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:]
	for i := 0; i < len(arg); i++ {
		for j := i; j < len(arg); j++ {
			if arg[i] > arg[j] {
				arg[i], arg[j] = arg[j], arg[i]
			}
		}
	}
	for _, a := range arg {
		for _, p := range a {
			z01.PrintRune(p)
		}
		z01.PrintRune('\n')
	}
}
