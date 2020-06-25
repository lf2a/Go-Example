package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func leitura() {
	arquivo := "arquivo.txt"

	// irá retornar o conteudo do arquivo em um []byte
	buffer, err := ioutil.ReadFile(arquivo)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// convertendo um []byte para string
	texto := string(buffer)
	fmt.Println(texto)

	data := bufio.NewScanner(strings.NewReader(texto))
	data.Split(bufio.ScanRunes)

	// le caractere por caractere
	for data.Scan() {
		fmt.Print(data.Text())
	}

	fmt.Println("\nFim.")
}
