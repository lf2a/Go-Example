package main

import (
	"fmt"
	"path"
)

func base() {
	// retorna o ultimo elemento do caminho
	fmt.Println(path.Base("/b"))
	fmt.Println(path.Base("/"))
	fmt.Println(path.Base(""))
}

func clear() {
	paths := []string{
		"a/c",
		"a//c",
		"a/c/.",
		"a/c/b/..",
		"/../a/c",
		"/../a/b/../././/c",
		"",
	}

	for _, p := range paths {
		fmt.Printf("Clean(%q) = %q\n", p, path.Clean(p))
	}
}

func main() {
	// base()
	clear()

	fmt.Println(path.Ext("/a/b/c/bar.css"))
	fmt.Println(path.Join("a", "b", "c"))
	fmt.Println(path.Join("a", "b/c"))
	fmt.Println(path.Join("a/b", "c"))
	fmt.Println(path.Join("", ""))
	fmt.Println(path.Join("a", ""))
	fmt.Println(path.Join("", "a"))
}
