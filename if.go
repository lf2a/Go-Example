package main

import "fmt"

func main() {
	// exemplo de if, else if e else
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println("even")
		} else if i%3 == 0 {
			fmt.Println(i)
		} else {
			fmt.Println("odd")
		}
	}

	fmt.Println("")
	for j := 0; j < 100; j++ {
		// atribuicao de valor a uma variavel dentro da expressao if
		if p := j * 10; p < 50 {
			fmt.Println(p)
		}
	}
}
