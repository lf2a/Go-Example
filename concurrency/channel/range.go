package main

import "fmt"

func testRange(c chan int) {
	/*
		Somente o remetente deve fechar um canal, nunca o receptor.
		Enviar em um canal fechado causará pânico.
		Canais não são como arquivos; você geralmente não precisa fechá-los.
		O fechamento é necessário apenas quando o receptor deve ser informado
		de que não há mais valores chegando, como para finalizar um loop de faixa.
	*/
	defer close(c)

	for i := 0; i < 5; i++ {
		c <- i
	}
}

func exemploRange() {
	c := make(chan int, 5)
	go testRange(c)

	// O loop `for i: = range c` recebe valores do canal repetidamente até que seja fechado.
	for i := range c {
		fmt.Println("exemploRange():", i)
	}
}
