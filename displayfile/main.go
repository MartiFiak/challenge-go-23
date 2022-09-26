package main

import (
	"fmt"
	"os"
	"io/ioutil"
)

func main () {
	FileName := "quest8.txt"

	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}
	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	} 
	data, err := ioutil.ReadFile(FileName)

	if err != nil {
		fmt.Println("Almost there!!")
		return
	}
	fmt.Println(data)
}
