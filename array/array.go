package main

import "fmt"

// len(): é o total de elementos que há dentro do array.
// cap(): é a quantidade que o array suporta, quando essa capacidade inicial
// é atingida ele dobra sua capacidade e armazena em um novo endereco na memoria

/*	Funciona tambem com estes tipos
- bool
- int  int8  int16  int32  int64
- uint uint8 uint16 uint32 uint64 uintptr
- byte -> alias for `uint8`
- rune -> alias for `int32` represents a Unicode code point
- float32 float64
- complex64 complex128
*/

func main() {
	var nome [10]string
	fmt.Printf("1 - var nome []string: %v\n", nome)
	fmt.Printf("1 - len: %d\tcap: %d\n", len(nome), cap(nome))

	nome[0] = "Luiz"
	nome[2] = "Filipy"
	fmt.Printf("2 - var nome []string: %v\n", nome)
	fmt.Printf("2 - len: %d\tcap: %d\n", len(nome), cap(nome))

	paises := []string{"Brasil", "EUA", "Canada"}
	fmt.Printf("3 - paises: %v\n", paises)
	fmt.Printf("3 - len: %d\tcap: %d\n", len(paises), cap(paises))

	doisPaises := paises[:2]
	fmt.Printf("4 - doisPaises: %v\n", doisPaises)
	fmt.Printf("4 - len: %d\tcap: %d\n", len(doisPaises), cap(doisPaises))

	doisPaises2 := paises[1:]
	fmt.Printf("5 - doisPaises: %v\n", doisPaises2)
	fmt.Printf("5 - len: %d\tcap: %d\n", len(doisPaises2), cap(doisPaises2))

	umPais := paises[1:2]
	fmt.Printf("6 - umPais: %v\n", umPais)
	fmt.Printf("6 - len: %d\tcap: %d\n", len(doisPaises2), cap(doisPaises2))

	cidades := make([]string, 10)
	cidades[0] = "Belém"
	cidades[1] = "São Paulo"
	cidades[2] = "Rio de Janeiro"
	cidades[3] = "New York"
	cidades[4] = "Boston"
	fmt.Printf("7 - cidades: %v\n", cidades)
	fmt.Printf("7 - len: %d\tcap: %d\n", len(cidades), cap(cidades))

	cidades = append(cidades, "São Francisco", "California", "Ananindeua")
	fmt.Printf("8 - cidades: %v\n", cidades)
	fmt.Printf("8 - len: %d\tcap: %d\n", len(cidades), cap(cidades))

}
