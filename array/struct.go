package main

import "fmt"

type Pessoa struct {
	Nome  string
	Idade byte
}

func main() {
	p := []Pessoa{
		{Nome: "Luiz Filipy", Idade: 20},
		{Nome: "Ana Paula", Idade: 19},
		{Nome: "Larissa Fernandes", Idade: 23},
	}

	fmt.Printf("1 - []Pessoa: %v\n", p)
	fmt.Printf("1 - len: %d\tcap: %d\n", len(p), cap(p))

	var p2 []Pessoa
	p2 = append(p2, p[0])
	p2 = append(p2, p[1])
	p2 = append(p2, p[2])
	p2 = append(p2, Pessoa{Nome: "Jose Silva", Idade: 30})
	p2 = append(p2, Pessoa{Nome: "Leticia Souza", Idade: 21})
	fmt.Printf("2 - var nome []Pessoa: %v\n", p2)
	fmt.Printf("2 - len: %d\tcap: %d\n", len(p2), cap(p2))

	p2 = p2[2:4]
	fmt.Printf("2 - []Pessoa: %v\n", p2)
	fmt.Printf("2 - len: %d\tcap: %d\n", len(p2), cap(p2))

	p3 := []struct {
		name string
		age  uint8
	}{
		{"Luiz", 45},
		{"Carla", 19},
		{"Pedro", 23},
		{"Lua", 64},
		{"José", 34},
		{"Maria", 12},
	}
	fmt.Printf("3 - []struct: %v\n", p3)
	fmt.Printf("3 - len: %d\tcap: %d\n", len(p3), cap(p3))

	type Marca struct {
		Marca string
	}

	marcas := make([]Marca, 10)
	marcas[0] = Marca{"Samsung"}
	marcas[1] = Marca{"Xiaomi"}
	marcas[2] = Marca{"Apple"}
	fmt.Printf("4 - []Marca: %v\n", p3)
	fmt.Printf("4 - len: %d\tcap: %d\n", len(p3), cap(p3))
}
