package main

import (
	"log"
	"os"
)

func renomear() {
	original := "arquivo.txt"
	novo := "arquivo1.txt"

	// pode ser usado para renomear arquivo e/ou mover
	err := os.Rename(original, novo)
	if err != nil {
		log.Fatal(err)
	}
}
