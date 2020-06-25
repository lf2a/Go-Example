package main

import "fmt"

// funcao com 2 parametros do mesmo tipo e apenas um retorno
// se os parametros forem do mesmo tipo pode omitir a declaração do tipo
// ex: add(v1, v2, v3 int) ou funcao(i int, nome, sobrenome string)
func add(v1, v2 int) int {
	return v1 + v2
}

// funcao que retorna multiplos valores
func name(age int, name string) (int, string) {
	return age, name
}

// funcao que retorna valores nomeados
// ira retornar os valores das variaveis: "name" e "age"
func person(n string, a int) (name string, age int) {
	name = "Brasil " + n
	age = a
	return
}

// funcao que recebe como argumento uma funcao e retorna um inteiro
func calc(fn func(v1, v2 int) int) int {
	return fn(3, 4)
}

// funcao que retorna uma funcao com parametro do tipo float64 (double)
func test1() func(float64) float64 {
	return func(y float64) float64 {
		return y / 2
	}
}

// funcao com 2 parametros o primeiro recebe um inteiro
// e o segundo recebe uma quantidade indefinida de inteiros
func variadic(num int, nums ...int) {
	fmt.Println("funcao 6:", num, nums)
}

func main() {
	fmt.Println("funcao 1:", add(23, 2))

	a, n := name(19, "Luiz")
	fmt.Println("funcao 2:", a, n)

	x, y := person("Luiz", 19)
	fmt.Println("funcao 3:", x, y)

	// funcao anonima
	sum := func(v1, v2 int) int {
		return v1 + v2
	}
	fmt.Println("funcao 4:", calc(sum))
	fmt.Println("funcao anonima:", sum(5, 9))

	t1 := test1()
	fmt.Println("funcao 5", t1(12))

	variadic(12, 32, 65, 56, 7, 89)
}
