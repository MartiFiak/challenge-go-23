package main
import (
    "os"
    "github.com/01-edu/z01"
)
func main() {
    arr := []rune(os.Args[0]) // Conversion de Args[0] (qui contient le programmname) en tableau de rune
    for i := 2; i < len(arr); i++ { // On parcourt le tableau à partir de la seconde position pour esquiver le ./
        z01.PrintRune(arr[i])
    }
    z01.PrintRune('\n')
}
