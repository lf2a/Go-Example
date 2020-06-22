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

func writer() {
	w := bufio.NewWriter(os.Stdout)

	// Flush grava todos os dados em buffer no io.Writer subjacente.
	defer w.Flush()

	// Write escreve o conteúdo do []byte no buffer. Retorna o número de bytes gravados.
	w.Write([]byte("teste\n"))

	// WriteString grava uma string. Retorna o número de bytes gravados.
	w.WriteString("LUIZ FILIPY\n")

	fmt.Fprint(w, "Hello, ")
	fmt.Fprint(w, "world!\n")

	// retorna quantos bytes não são usados no buffer.
	fmt.Println("A:", w.Available())

	// retorna o número de bytes que foram gravados no buffer atual.
	fmt.Println("B:", w.Buffered())

	// retorna o tamanho do buffer subjacente em bytes.
	fmt.Println(w.Size())
}
