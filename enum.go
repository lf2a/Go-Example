package main

import "fmt"

type Semana uint8

const (
	domingo Semana = iota + 1
	segunda        // 2
	terca          // 3
	quarta         // 4
	quinta         // 5
	sexta          // 6
	sabado         // 7
)

func enum() {
	fmt.Println(domingo)
}
