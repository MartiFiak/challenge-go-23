package PrintProgramName

import (
	"os"

	"github.com/01-edu/z01"
)

func PrintProgramName() {
	pr, _ := os.Executable()
	z01.PrintRune(" %s\n", pr)
	z01.PrintRune(" %s\n", os.Args[0])
}
