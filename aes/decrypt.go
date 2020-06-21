package main

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"io/ioutil"
)

func decrypt() {
	chave := []byte("estasenhaprecisater32bytes!!!!!!")
	ciphertext, err := ioutil.ReadFile("crypt.txt")
	if err != nil {
		fmt.Println(err)
	}

	c, err := aes.NewCipher(chave)
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
