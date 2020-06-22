package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	O pacote bufio implementa I/O em buffer. Ele envolve um objeto io.Reader ou io.Writer,
	criando outro objeto (Reader ou Writer) que também implementa a interface, mas fornece
	buffer e ajuda para I/O de texto.
*/

func read() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("> ")

	// A leitura avança o Scanner para o próximo token, que estará disponível pelo método
	// Bytes() ou Text(). Retorna false quando a verificação é interrompida, atingindo o
	// final da entrada ou um erro. Após a varredura retornar falsa, o método Err retornará
	// qualquer erro que tenha ocorrido durante a varredura, exceto que, se fosse io.EOF,
	// Err retornará nulo. O scanner entra em pânico se a função de divisão retornar muitos
	// tokens vazios sem avançar a entrada. Este é um modo de erro comum para scanners.
	_ = scanner.Scan()

	// retorna (em texto) o token mais recente gerado por uma chamada para o Scan()
	fmt.Println("> ", scanner.Text())

	// retorna (em bytes) o token mais recente gerado por uma chamada para o scan.
	fmt.Println("bytes: ", scanner.Bytes())

	// retorna o primeiro erro não-EOF encontrado pelo Scanner.
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "erro ao ler:", err)
	}
}
