package main

/*
Maps são tipos de dados associativos padrões no Go (às vezes chamados de hashes ou
dicionários em outras linguagens).
*/

import "fmt"

type Aviao struct {
	Ano int
}

func main() {
	// criando um map de struct aviao
	var avioes map[string]Aviao

	avioes = make(map[string]Aviao)
	avioes["novo"] = Aviao{1999}
	avioes["antigo"] = Aviao{2000}

	fmt.Println(avioes["novo"])
	fmt.Println(avioes)

	// apagando a chave "antigo"
	delete(avioes, "antigo")
	fmt.Println(avioes)

	// verificando se existe, se existir a variavel "ok" ira retornar true
	aviao, ok := avioes["antigo"]
	fmt.Println(aviao, ok)

	// ==========

	var avioes2 = map[string]Aviao{
		"luiz":   {1956},
		"filipy": {1890},
	}

	fmt.Println(avioes2)

	// ==========

	// chave string, valor int
	nums := make(map[string]int)
	nums["a"] = 12
	nums["b"] = 15

	fmt.Println(nums)
	fmt.Println(nums["b"])

	nums2 := map[string]int{"c": 23, "d": 34}
	fmt.Println(nums2)
}
