package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:] // trieur
	for i := 0; i < len(arg); i++ { // définit position i(index) qui parcourt arg
		for j := i; j < len(arg); j++ { // définit position j qui parcourt arg
			if arg[i] > arg[j] { 
				arg[i], arg[j] = arg[j], arg[i] // inversement de position entre i et j
			}
		}
	}
	for _, a := range arg {
		for _, p := range a {
			z01.PrintRune(p) //print
		}
		z01.PrintRune('\n')
	}
}
