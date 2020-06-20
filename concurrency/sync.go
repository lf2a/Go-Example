package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	Para esperar que várias goroutines terminem, podemos usar um WaitGroup
*/

// é usado para esperar o programa terminar a execução das goroutines

func loop(s string, wg *sync.WaitGroup) {
	// um WaitGroup deve ser passado para funções por ponteiro.
	defer wg.Done()

	for i := 0; i <= 2; i++ {
		time.Sleep(time.Second * time.Duration(2))
		fmt.Printf("%s: %d\n", s, i)
	}
}

func syncEcemplo() {
	var wg sync.WaitGroup
	fmt.Println("Inicio")

	/*
		contador: se o contador chegar a zero todas as goroutines bloqueadas
		na espera são liberadas. Se o contador ficar negativo irá ser lançado `panic`
	*/
	wg.Add(2)

	go loop("l1", &wg)
	go loop("l2", &wg)

	fmt.Println("Esperando terminar")
	wg.Wait()
	fmt.Println("Pronto")
}
