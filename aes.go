package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
)

func Crypt() {
	text := []byte("My Super Secret Code Stuff")
	key := []byte("passphrasewhichneedstobe32bytes!")

	fmt.Printf("%s | %s\n", key, text)

	// gera uma nova cifra aes usando uma chave de 32 bytes
	c, err := aes.NewCipher(key)

	if err != nil {
		fmt.Println(err)
	}

	/*
		gcm or Galois/Counter Mode, é um modo de operação para cifras
		de bloco criptográficas de chave simétrica
	*/
	gcm, err := cipher.NewGCM(c)

	if err != nil {
		fmt.Println(err)
	}

	/*
		cria uma nova matriz de bytes do tamanho do
		nonce que deve ser passado para o Seal
	*/
	nonce := make([]byte, gcm.NonceSize())

	/*
		preenche nosso nonce com uma sequência
		aleatória criptograficamente segura
	*/
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		fmt.Println(err)
	}

	/*
		aqui criptografamos nosso texto usando a função Seal
		Seal criptografa e autentica texto sem formatação,
		autentica os dados adicionais e anexa o resultado
		ao dst, retornando a atualização do slice.
		O nonce deve ter bytes NonceSize() e ser exclusivo
		para todos os tempos, para uma determinada chave.
	*/
	b := gcm.Seal(nonce, nonce, text, nil)

	ioutil.WriteFile("my.data", b, 0644)

	fmt.Println(b)
}

func Decrypt() {
	key := []byte("passphrasewhichneedstobe32bytes!")
	ciphertext, err := ioutil.ReadFile("my.data")
	if err != nil {
		fmt.Println(err)
	}

	c, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		fmt.Println(err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		fmt.Println(err)
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(plaintext))
}
