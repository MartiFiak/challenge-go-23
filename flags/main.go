package main

import (
	"fmt"
	"os"
)

func main() {
	command := os.Args
	GetFlags(command)
}

func Insert(strOrigine string, strToAdd string) string {
	return strOrigine + strToAdd // Problème avec le test go run . "-i=Q 0d 6JI S 4j" "lsJ  FhbrkYGN"
}

func Order(strToSort string) string {
	runeSort := []rune(strToSort)
	for i := 0; i < len(runeSort); i++ {
		for j := i + 1; j < len(runeSort); j++ {
			if runeSort[i] > runeSort[j] {
				runeSort[i], runeSort[j] = runeSort[j], runeSort[i]
			}
		}
	}
	c := ""
	for _, i := range runeSort {
		c += string(i)
	}
	return c
}

func Help() {
	fmt.Println("--insert")
	fmt.Println("  -i")
	fmt.Println("     This flag inserts the string into the string passed as argument.")
	fmt.Println("--order")
	fmt.Println("  -o")
	fmt.Println("     This flag will behave like a boolean, if it is called it will order the argument.")
}
func GetFlags(command []string) {
	if len(command) == 1 {
		Help()
	} else {
		if command[1] == "--help" || command[1] == "-h" {
			Help()
		} else if len(command[1]) > 8 {
			if command[1][:8] == "--insert" {
				if command[2] == "--order" || command[2] == "-o" {
					fmt.Println(Order(Insert(command[3], command[1][9:])))
				} else {
					fmt.Println(Insert(command[2], command[1][9:]))
				}
			} else if command[1][:2] == "-i" {
				if command[2] == "--order" || command[2] == "-o" {
					fmt.Println(Order(Insert(command[3], command[1][3:])))
				} else {
					fmt.Println(Insert(command[2], command[1][3:]))
				}
			}
		} else if command[1][:2] == "-i" {
			if command[2] == "--order" || command[2] == "-o" {
				fmt.Println(Order(Insert(command[3], command[1][3:])))
			} else {
				fmt.Println(Insert(command[2], command[1][3:]))
			}
		} else if command[1] == "--order" || command[1] == "-o" {
			fmt.Println(Order(command[2]))
		} else {
			fmt.Println(command[1])
		}
	}
}
