package main

import (
	"fmt"
	"log"
	"os"
)

func info() {
	// usado para obter informações de arquivos ou diretorios
	arquivo, err := os.Stat("obter_informacoes.go")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Nome:", arquivo.Name())           // nome base do arquivo
	fmt.Println("Tamanho:", arquivo.Size())        // tamanho em bytes
	fmt.Println("Permissão:", arquivo.Mode())      // modo do arquivo
	fmt.Println("Modificado:", arquivo.ModTime())  // ultima modificação
	fmt.Println("É diretorio?: ", arquivo.IsDir()) // abreviação para Mode().IsDir()
}
