package main

import (
	"fmt"
)

func main() {
	aSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(aSlice)
	fmt.Println("---------------")

	integer := make([]int, 20)

	integer = append(integer, 123456)

	for i := 0; i < len(integer); i++ {
		fmt.Print(integer[i], " ")
	}

	fmt.Println()
	fmt.Println("---------------")
	fmt.Println(integer[0], integer[20])
	fmt.Println("---------------")
	fmt.Println(integer[len(integer)-1])

	s2 := integer[19:21]
	fmt.Println("---", s2, "---")

	fmt.Println()
	fmt.Println("=====================")

	s1 := make([]int, 5)
	reSlice := s1[1:3]
	fmt.Println(s1)
	fmt.Println(reSlice)
	reSlice[0] = -100
	reSlice[1] = 123456
	fmt.Println(s1)
	fmt.Println(reSlice)
}
