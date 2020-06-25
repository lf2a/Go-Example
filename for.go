package main

import "fmt"

func main() {
	// for padão
	for i := 0; i < 5; i++ {
		fmt.Printf("index: %d\n", i)
	}

	fmt.Println("===================================")

	// `while`
	i := 0
	for i < 5 {
		fmt.Printf("i %d\n", i)
		i++
	}

	// loop infinito
	// for {}

	// range (foreach)
	w := []int{1, 2, 4, 8}
	// o index é o primeiro valor (i) e o valor do indice é (v)
	for i, v := range w {
		fmt.Printf("index: %d value: %d\n", i, v)
	}

	// `_` omitindo (ignorando) o valor
	for _, v := range w {
		fmt.Printf("valor: %d\n", v)
	}
}
