package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func write() {
	// cria um buffer para gravar o arquivo
	// buf := new(bytes.Buffer)
	file, err := os.Create("arquivo.zip")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	// cria um arquivo zip
	zipFile := zip.NewWriter(file)

	// alguns arquivos para `zipar`
	var files = []struct{ Name, Content string }{
		{"file1.txt", "file 1\ntxt example"},
		{"file2.txt", "file 2\ntxt example"},
		{"file3.txt", "file 3\ntxt example"},
		{"file4.txt", "file 4\ntxt example"},
	}

	for _, file := range files {
		// criar um arquivo
		f, err := zipFile.Create(file.Name)

		if err != nil {
			fmt.Println(err)
		}

		// escreve o conteudo no arquivo
		_, err = f.Write([]byte(file.Content))

		if err != nil {
			fmt.Println(err)
		}
	}

	ff, _ := zipFile.Create("array.go")
	data, _ := ioutil.ReadFile("array.go") // data []byte
	ff.Write(data)

	// termina a gravação do arquivo zip, gravando o diretório central
	defer zipFile.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func read() {
	// abre o arquivo zip
	r, err := zip.OpenReader("arquivo.zip")
	if err != nil {
		fmt.Println(err)
	}
	defer r.Close()

	// percorre todos os arquivo no zip
	for _, f := range r.File {
		fmt.Printf("Arquivo: %s\n", f.Name)
		rc, err := f.Open()
		if err != nil {
			fmt.Println(err)
		}

		_, err = io.CopyN(os.Stdout, rc, 3000)
		if err != nil {
			fmt.Println(err)
		}
		defer rc.Close()
	}
}

func descomprimir() {
	// abre arquivo .zip
	zipFile, err := zip.OpenReader("arquivo.zip")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer zipFile.Close()

	for i, file := range zipFile.File {

		fmt.Printf("descomprimindo arquivo #%02d %v\n", i+1, file.Name)

		// abre reader para ler arquivo de dentro do zip
		reader, err := file.Open()
		if err != nil {
			fmt.Println(err)
			return
		}
		defer reader.Close()

		var f *os.File
		// abre arquivo de destino
		f, err = os.OpenFile(
			file.Name,
			os.O_WRONLY|os.O_CREATE|os.O_TRUNC,
			file.Mode())
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		// grava arquivo de destino
		_, err = io.Copy(f, reader)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func ZipTest() {
	fmt.Println("Teste")

	write()
	read()
	descomprimir()
}
