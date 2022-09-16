package printprogramname

import (
	"fmt"
	"github.com/01-edu/z01.PrintRune"
	"os"
)

func printprogramname() {
	arguments := os.Args

	z01.PrintRune("Number of arguments: %v\n", len(arguments))

	if len(arguments) >= 3 {
		z01.PrintRune("argument# 3 is: %v", os.Args[2])
	}
}
