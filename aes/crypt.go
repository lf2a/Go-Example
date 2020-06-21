package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
)

// Exemplo retirado do site tutorialedge.net: https://tutorialedge.net/golang/go-encrypt-decrypt-aes-tutorial/

func crypt() {
	texto := []byte("Luiz Filipy - Brasil 20 1999")
	chave := []byte("estasenhaprecisater32bytes!!!!!!")

	fmt.Println("Texto:", texto)
	fmt.Println("Senha:", chave)

	// gera uma nova cifra aes usando uma chave de 32 bytes
	c, err := aes.NewCipher(chave)

	if err != nil {
		fmt.Println(err)
	}

	// gcm or Galois/Counter Mode, é um modo de operação para cifras de bloco criptográficas de chave simétrica
	gcm, err := cipher.NewGCM(c)

	if err != nil {
		fmt.Println(err)
	}

	// cria uma nova matriz de bytes do tamanho do nonce que deve ser passado para o Seal
	nonce := make([]byte, gcm.NonceSize())

	// preenche nosso nonce com uma sequência aleatória criptograficamente segura
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	/*
		aqui criptografamos nosso texto usando a função Seal Seal criptografa e autentica texto sem formatação,
		autentica os dados adicionais e anexa o resultado ao dst, retornando a atualização do slice.
		O nonce deve ter bytes NonceSize() e ser exclusivo para todos os tempos, para uma determinada chave.
	*/
	b := gcm.Seal(nonce, nonce, texto, nil)

	ioutil.WriteFile("crypt.txt", b, 0644)

	fmt.Println(b)
}
