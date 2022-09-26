package main

import "github.com/01-edu/z01"

func main() {
	p := "x = " + IntegerToString(42) + ", y = " + IntegerToString(21) + "\n"
	for _, c := range p {
		z01.PrintRune(c)
	}
}

func IntegerToString(nbr int) string {
	snbr := ""
	var allunit []int
	if nbr == 0 {
		return "0"
	} else {
		for nbr > 0 {
			allunit = append(allunit, nbr%10)
			nbr = nbr / 10
		}
	}
	for c := len(allunit) - 1; c >= 0; c-- {
		snbr += string(rune(allunit[c] + 48))
	}
	return snbr
}
