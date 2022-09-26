package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	} else if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	} else {
		file, _ := os.Open(os.Args[1])
		b, _ := ioutil.ReadAll(file)
		fmt.Print(string(b))
	}
}
