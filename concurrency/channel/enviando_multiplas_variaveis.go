package main

import "fmt"

func funcaoComErro(f chan func() (int, error)) {
	f <- func() (int, error) { return 99, nil }
}

type Pessoa struct {
	Idade int8
	Nome  string
}

func funcaoComEstrutura(p chan Pessoa) {
	p <- Pessoa{20, "Luiz Filipy"}
}

func enviandoMultiplasVariaveis() {
	c1 := make(chan func() (int, error))
	go funcaoComErro(c1)

	res, err := (<-c1)()
	if err == nil {
		fmt.Println("funcaoComErro():", res)
	} else {
		fmt.Println("funcaoComErro():", err)
	}

	// cria um canal do tipo `Pessoa`
	c2 := make(chan Pessoa)
	go funcaoComEstrutura(c2)

	res2 := <-c2
	fmt.Println("Pessoa:", res2)

	fmt.Println("FIM")
}
