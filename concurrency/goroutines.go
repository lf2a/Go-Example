package main

/*
	As goroutines são muito menores do que os threads,
	normalmente levam cerca de 2kB de espaço de pilha
	para inicializar, em comparação com um thread de 1Mb.
	As goroutines são executadas no mesmo espaço de endereço,
	portanto, o acesso à memória compartilhada deve ser sincronizado.
*/

import (
	"fmt"
	"time"
)

func test() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(1000 * time.Millisecond)
	}
}

func goroutine_exemplo() {
	// para invocar esta função em uma goroutine é necessário por a palavra chage `go` na frente da chamada
	// Estas goroutines serão executadas simultaneamente
	go test()
	go test()

	go func() {
		// iniciando uma goroutine através de uma função anonima
		fmt.Println("Função anonima")
	}()

	// executando de forma sincrona
	test()
}
