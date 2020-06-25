package main

import (
	"log"
	"os"
)

func write() {
	arquivo, err := os.Create("arquivo1.txt")
	defer arquivo.Close()

	if err != nil {
		log.Fatal(err)
	}

	arquivo.Write([]byte("luiz Filipy")) // para escrever algo no arquivo
	arquivo.WriteString("\n\nBrasil 20")
}
