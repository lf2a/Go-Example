package main

import (
	"fmt"
	"log"
	"os"
)

func diretorio() {
	_, err := os.Stat("novo_dir")

	if os.IsNotExist(err) {
		errDir := os.MkdirAll("novo_dir", 0755)
		if errDir != nil {
			log.Fatal(err)
		}
		fmt.Println("Diretorio Criado.")

	}
}
