package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 100; i++ {

		if i%20 == 0 {
			continue
		}

		if i == 95 {
			break
		}

		fmt.Print(i, " ")
	}

	fmt.Println("------------ exit first loop ------------")

	i := 10
	for {
		if i < 0 {
			break
		}
		fmt.Print(i, " ")
		i--
	}
	fmt.Println("------------ exit secondary loop ------------")

	i = 0
	anExpression := true
	for ok := true; ok; ok = anExpression {
		if i > 10 {
			anExpression = false
		}

		fmt.Print(i, " ")
		i++
	}
	fmt.Println("------------ exit three loop ------------")

	anArray := [5]int{0, 1, -1, 2, -2}
	for i, value := range anArray {
		fmt.Println("index: ", i, "value: ", value)
	}
	fmt.Println("------------ exit four loop ------------")
	fmt.Println(anArray)
	fmt.Println(len(anArray))
}
