package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func SHATest() {
	// sha 256
	sum := sha256.Sum256([]byte("hello world\n"))
	fmt.Printf("%x\n", sum)

	// sha 512
	sum1 := sha512.Sum512([]byte("hello world\n"))
	fmt.Printf("%x\n", sum1)
}
