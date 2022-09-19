package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	for _, t := range os.Args[1:] {
		if t == "--upper" {
			for _, csc := range os.Args[2:] {
				z01.PrintRune(rune(Atoi(csc) + 64))
			}
			break
		} else if Atoi(t) > 0 && Atoi(t) <= 26 {
			z01.PrintRune(rune(Atoi(t) + 96))
		} else {
			z01.PrintRune(32)
		}
	}
	if len(os.Args) > 1 {
		z01.PrintRune('\n')
	} else {
		return
	}
}

func Atoi(s string) int {
	nb := 0
	count := 0
	neg := false
	for i, r := range s {
		if r == 43 && i == 0 {
			count += 1
		} else if r == 45 && i == 0 {
			neg = true
			count += 1
		} else if r-48 < 0 || r-48 > 9 {
			return 0
		} else if count > 1 {
			return 0
		} else {
			nb = nb*10 + int(r-48)
		}
	}
	if neg {
		nb = -nb
	}
	return nb
}
