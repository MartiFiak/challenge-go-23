package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) == 1 {
		aze, _ := ioutil.ReadAll(os.Stdin)
		ok(string(aze))
	}
	tab := []string(os.Args[1:])
	if len(os.Args) > 1 {
		data, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			ok(string("ERROR: " + err.Error()))
			z01.PrintRune(('\n'))
			os.Exit(1)
			ok(string(data))
		} else {
			for _, j := range tab {
				data1, err := ioutil.ReadFile(j)
				if err != nil {
					ok(`ERROR: ` + err.Error())
					z01.PrintRune(('\n'))
					os.Exit(1)
				} else {
					ok(string(data1))
				}
			}
		}
	}
}

func ok(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}
