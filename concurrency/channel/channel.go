package main

import (
	"fmt"
	"time"
)

/*
	Channel: serve para Goroutines se comunicarem entre si.

	data := <- a // lê do canal a
	a <- data // escreve no canal a

	Os envios e recebimentos para um canal são bloqueando por padrão
	Por padrão, envia e recebe o bloco até que o outro lado esteja pronto.
	Isso permite que as goroutines sejam sincronizadas sem bloqueios explícitos ou variáveis de condição.
*/

func test(c chan string, s string, segundos int) {
	time.Sleep(time.Duration(segundos) * time.Second)
	c <- s
	fmt.Printf("test(): %s\n", s)
}

func channel_exemplo() {
	fmt.Println("INICIO")

	c1 := make(chan string) // não buferizado
	c2 := make(chan string)

	// Canais podem ser armazenados em buffer.
	// Forneça o tamanho do buffer como o segundo argumento a
	// ser feito para inicializar um canal em buffer:
	//c := make(chan string, 2) // buferizado

	// quando sair da função ele irão ser fechadas
	defer close(c1)
	defer close(c2)

	go test(c1, "valor1", 5)
	valor1 := <-c1

	// passando o retorno a channel `c1` como valor para a segunda goroutine
	// irá esperar `valor1` quando estiver pronto
	go test(c2, valor1+" + valor2", 1)
	valor2 := <-c2

	fmt.Printf("FIM: %s - %s\n", valor1, valor2)
}
