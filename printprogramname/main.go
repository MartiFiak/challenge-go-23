package PrintProgramName

import (
	"os"
	"github.com/01-edu/z01.PrintRune"
)

func PrintProgramName() {
	programName := os.Args[0]
	z01.PrintRune(programName)
}
