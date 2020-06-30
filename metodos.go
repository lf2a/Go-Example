package main

import "fmt"

type Cachorro struct {
	Nome string
	Raca string
}

func (c *Cachorro) getInfo() {
	fmt.Printf("Método: %s é da raça %s\n", c.Nome, c.Raca)
}

func getInfo2(c *Cachorro) {
	fmt.Printf("Função que recebe o objeto da struct por parametro: %s é da raça %s\n", c.Nome, c.Raca)
}

type grandeInt int64

func (b *grandeInt) print() {
	fmt.Println(*b, b)
}

func main() {
	p1 := &Cachorro{"Doguinho", "Pit Bull"}
	// metodo
	p1.getInfo()

	// equivalente ao de cima
	getInfo2(p1)

	// outro exemplo
	b := grandeInt(1234567890)
	b.print()
}
