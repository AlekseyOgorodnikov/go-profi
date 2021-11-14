package main

import "fmt"

type aStructure struct {
	person string
	height int
	weight int
}

func main() {
	var s1 aStructure

	p1 := aStructure{"smt", 111, 20}
	fmt.Println(s1)

	fmt.Println(p1)

	g1 := aStructure{height: 20, weight: 22, person: "sadasd"}
	fmt.Println(g1)

	a1 := aStructure{height: 222, weight: 222222554}
	fmt.Println(a1.person)
}
