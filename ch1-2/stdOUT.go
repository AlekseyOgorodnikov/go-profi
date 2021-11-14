package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	myString := ""
	arguments := os.Args
	if len(arguments) == 1 {
		myString = "Please give me one argument!"
	} else {
		myString = arguments[1]
	}

	io.WriteString(os.Stdout, myString)
	io.WriteString(os.Stdout, "\n")

	prints()
}

func prints() {
	i, k := 2, 3
	j, k := 4, 5
	fmt.Println(i, j, k)
}
