package main

import (
	"fmt"
	"os"
	"time"
)

/*
	A instrução select permite que uma goroutine aguarde várias operações de comunicação.
	Um select bloqueia até que um de seus casos possa ser executado e, em seguida,
	executa esse caso. Ele escolhe um aleatoriamente se vários estiverem prontos.

	Os timeouts são importantes para programas que se conectam a recursos externos ou que,
	de outra forma, precisam limitar o tempo de execução. A implementação de tempos limite
	no Go é fácil e elegante, graças aos channels e select.
*/

func main() {
	c1 := make(chan string, 2)
	c2 := make(chan string, 2)
	c3 := make(chan string, 2)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "1()"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c2 <- "2()"
	}()

	go func() {
		time.Sleep(4 * time.Second)
		c3 <- "ok"
		close(c3)
	}()

	for ; ; {
		select {
		case msg1 := <-c1:
			fmt.Println("1", msg1)
		case msg2 := <-c2:
			fmt.Println("2", msg2)
		case msg3, mais := <-c3:
			if mais {
				fmt.Println("ainda não", msg3)
			} else {
				fmt.Println("Saindo Agora!", msg3)
				os.Exit(0)
			}
			//default:
			// O caso default em uma seleção é executado se nenhum outro caso estiver pronto.
			//fmt.Println("default")
		}
	}
}
