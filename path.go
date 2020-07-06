package main

import (
	"fmt"
	"path"
)

func base() {
	// retorna o ultimo elemento do path
	fmt.Println(path.Base("/b"))
	fmt.Println(path.Base("/b/a"))
	fmt.Println(path.Base("/"))
	fmt.Println(path.Base(""))
}

func clean() {
	paths := []string{
		"a/c",
		"a//c",
		"a/c/.",
		"a/c/b/..",
		"/../a/c",
		"/../a/b/../././/c",
		"",
	}

	// Clean retorna o nome do caminho mais curto equivalente ao caminho por processamento puramente lexical.
	for _, p := range paths {
		fmt.Printf("Clean(%q) = %q\n", p, path.Clean(p))
	}
}

func main() {
	clean()
	base()

	// Ext retorna a extensão do nome do arquivo usada pelo path.
	fmt.Println(path.Ext("/a/b/c/bar.css"))

	// Join une qualquer número de elementos de path em um único path, separando-os com barras.
	// Elementos vazios são ignorados.
	fmt.Println(path.Join("a", "b", "c"))
	fmt.Println(path.Join("a", "b/c"))
	fmt.Println(path.Join("a/b", "c"))
	fmt.Println(path.Join("", ""))
	fmt.Println(path.Join("a", ""))
	fmt.Println(path.Join("", "a"))
}
