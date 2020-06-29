package main

import "log"

func init() {
	log.SetPrefix("LOG: ")
	// LOG: 0000/00/00 00:00:00 exemplo_simples.go:12: <mensagem>
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	log.Println("main started")

	log.Fatalln("fatal message")

	log.Panicln("panic message")
}
