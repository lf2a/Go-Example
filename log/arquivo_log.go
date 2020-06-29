package main

import (
	"log"
	"os"
)

func main() {
	arquivo, err := os.OpenFile(".log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	defer arquivo.Close()

	log.SetPrefix("LOG: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetOutput(arquivo)

	if err != nil {
		log.Println(err)
	}

	log.Println("text to append")
	log.Println("more text to append")
	log.Fatalln("fatal exemplo")
	log.Panicln("panic exemplo")
}
