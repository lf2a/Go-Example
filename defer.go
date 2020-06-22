package main

import "fmt"

/*
	Aviso:
	Se uma função adiada é avaliada como nula,
	a execução entra em pânico quando a função
	circundante termina e não quando a adiação é chamada.

	Não use defer em um loop, a menos que tenha certeza
	do que está fazendo.
*/

func ola() {
	defer fmt.Println("World")

	fmt.Println("Hello")
}

func main() {
	ola()

	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
