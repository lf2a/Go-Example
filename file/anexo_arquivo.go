package main

import (
	"fmt"
	"os"
)

func anexo() {
	conteudo := "\n\nAfinal"
	filename := "arquivo.txt"

	arquivo, err := os.OpenFile(filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	defer arquivo.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	arquivo.WriteString(conteudo)
}
