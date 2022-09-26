package main

import (
	"os"
	"io/ioutil"
	"fmt"
)

func readFromFille(f string) {	
	// read the whole content of file and pass it to file variable, in case of error pass it to err variable
	file, err := os.Open(f)
	if err != nil {
	   fmt.Printf("ERROR: open %s: no such file or directory \n", os.Args[1])
	   os.Exit(1)
	}
	// convert the file binary into a string using string
	fileContent := string(file)
	// print file content
	fmt.Println(fileContent)
 }
