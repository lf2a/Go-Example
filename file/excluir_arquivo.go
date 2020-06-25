package main

import (
	"log"
	"os"
)

func excluir() {
	err := os.Remove("arquivo1.txt")
	
	if err != nil {
		log.Fatal(err)
	}
}
