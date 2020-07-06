package main

import "fmt"

var a, b = 12, 19

func main() {
	fmt.Printf("antes da mudaça -> a: %d | b: %d\n\n",a, b)

	aa := &a
	bb := &b
	fmt.Printf("endereco -> aa: %d | bb:%d\n", aa, bb)
	fmt.Printf("valor -> *aa: %d | *bb:%d\n\n", *aa, *bb)

	*aa, *bb = 45, 78
	fmt.Printf("endereco depois da mudança -> aa: %d | bb:%d\n", aa, bb)
	fmt.Printf("valor depois da mudança -> *aa: %d | *bb:%d\n\n", *aa, *bb)

	fmt.Printf("depois da mudaça -> a: %d | b: %d\n",a, b)
}
