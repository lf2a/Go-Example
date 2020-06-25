package main

import (
	"io"
	"log"
	"os"
)

func copiar() {
	original, err := os.Open("arquivo1.txt")
	defer original.Close()

	if err != nil {
		log.Fatal(err)
	}

	novoArquivo, err := os.Create("arquivo.txt")
	defer novoArquivo.Close()

	if err != nil {
		log.Fatal(err)
	}

	copia, err := io.Copy(novoArquivo, original)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Bytes copiados %d.", copia)
}
