package main

import (
	"log"
	"os"
	"time"
)

// alerando permissao, propriedades e timestamps
func alterar() {
	// verificando se existe
	_, err := os.Stat("arquivo.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("Arquivo não existe.")
		}
	}

	log.Println("Arquivo existe.")

	// Mudando permissão wr-wr-wr.
	err = os.Chmod("arquivo.txt", 0666)
	if err != nil {
		log.Println(err)
	}

	// Alterar a propriedade do arquivo.
	err = os.Chown("arquivo.txt", os.Getuid(), os.Getgid())
	if err != nil {
		log.Println(err)
	}

	// Alterar timestamps de arquivo.
	// adicionar Um dia a partir de agora
	hora := time.Now().Add(24 * time.Hour)
	ultimoAcesso := hora
	ultimaModificacao := hora
	err = os.Chtimes("arquivo.txt", ultimoAcesso, ultimaModificacao)
	if err != nil {
		log.Println(err)
	}
}
