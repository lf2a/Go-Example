package main

import (
	"fmt"
	"os"
)

func main() {
	os.Setenv("NOME", "Luiz")
	os.Setenv("IDADE", "20")

	fmt.Printf("%s %s\n", os.Getenv("NOME"), os.Getenv("IDADE"))
}
