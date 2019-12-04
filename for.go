package main

import "fmt"

func ForTest() {
	for i := 0; i < 5; i++ {
		fmt.Printf("index: %d\n", i)
	}

	fmt.Println("")
	sub := 5
	for sub > 0 {
		sub--
		fmt.Printf("sub is %d\n", sub)
	}

	fmt.Println("")
	// `while`
	i := 0
	for i < 5 {
		fmt.Printf("i %d\n", i)
		i++
	}

	// infinity loop
	// for {}

	// range
	w := []int{1, 2, 4, 8}

	for i, v := range w {
		fmt.Printf("index: %d value: %d\n", i, v)
	}

	// `_` omitindo o valor
	for _, v := range w {
		fmt.Printf("valor: %d\n", v)
	}
}
