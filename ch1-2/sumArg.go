package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("Please more argument!")
		os.Exit(1)
	}

	// sumInt()
	sumFloat()
}

// sum arguments integer
func sumInt() {
	arguments := os.Args
	var sum int64

	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseInt(arguments[i], 10, 64)
		if err != nil {
			fmt.Println("Please input Integer numbers!")
			return
		}
		sum += n
	}
	fmt.Println(sum)
}

// sum arguments floating
func sumFloat() {
	arguments := os.Args
	var sum float64

	for i := 1; i < len(arguments); i++ {
		n, err := strconv.ParseFloat(arguments[i], 64)
		if err != nil {
			fmt.Println("Please input float numbers")
			return
		}
		sum += n
	}
	fmt.Println(sum)
}

// Напишите Go-программу, которая считывает целые числа до тех пор, пока
// не встретит во входных данных слово END.
