package main

import "fmt"

func main() {
	anArray := [5]int{0, 1, -1, 2, -2}
	for i, value := range anArray {
		fmt.Println("index: ", i, "value: ", value)
	}
	fmt.Println("------------ exit loop ------------")
	fmt.Println(anArray)
	fmt.Println(len(anArray))
	fmt.Println("------------ exit ------------")

	twoD := [4][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}, {13, 14, 15, 16}}
	threeD := [2][2][2]int{{{1, 0}, {-2, 4}}, {{5, -1}, {6, -5}}}
	fmt.Println("two array: ", twoD[0][2])
	fmt.Println("------------ exit ------------")

	for i := 0; i < len(threeD); i++ {
		v := threeD[i]
		for j := 0; j < len(v); j++ {
			m := v[j]
			for k := 0; k < len(m); k++ {
				fmt.Print(m[k], " ")
			}
		}
		fmt.Println()
	}

	for _, v := range threeD {
		for _, m := range v {
			for _, s := range m {
				fmt.Print(s, " ")
			}
		}
		fmt.Println("-------------------")
	}
}
