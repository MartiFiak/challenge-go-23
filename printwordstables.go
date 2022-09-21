package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, i := range a {
		srune := []rune(i)
		for _, j := range srune {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
