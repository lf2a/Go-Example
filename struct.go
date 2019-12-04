package main

import "fmt"

type personS struct {
	name string
	age  uint8
}

func StructTest() {
	fmt.Println(personS{"Luiz Filipy", 19})

	p1 := personS{"Luiz", 19}
	p1.name = "Filipy"
	fmt.Println(p1)

	luiz := &p1
	fmt.Println(luiz.name)
}
