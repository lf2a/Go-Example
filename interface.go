package main

import "fmt"

type Carro struct {
	modelo string
	marca  string
	preco  float32
}

type ICarro interface {
	info()
}

func (c Carro) info() {
	fmt.Printf("%s %s custa %6.2f\n", c.marca, c.modelo, c.preco)
}

func main() {
	// forma 1
	var i ICarro
	c1 := Carro{"Huayra", "Pagani", 1_500_00.00}

	i = &c1
	i.info()

	// forma 2
	var c2 ICarro = Carro{"Huracan", "Lamborghini", 250_000.00}
	c2.info()
}
