package piscine

import "github.com/01-edu/z01"

func PrintNbrInOrder(n int) {
	if n == 0 {
		z01.PrintRune(48)
	}
	var allunit []int
	if n < 0 {
		for n < 0 {
			allunit = append(allunit, n%10*(-1))
			n = n / 10
		}
	} else {
		for n > 0 {
			allunit = append(allunit, n%10)
			n = n / 10
		}
	}
	SortIntegerTable(allunit)
	for _, c := range allunit {
		z01.PrintRune(rune(c + 48))
	}
}
