package main

import "fmt"

type Semana uint8

const (
	domingo Semana = iota + 1
	segunda
	terca
	quarta
	quinta
	sexta
	sabado
)

func EnumTest() {
	fmt.Println(domingo)
}
