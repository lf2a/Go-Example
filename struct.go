package main

import "fmt"

type Person struct {
    name string
    age uint8
}

func main() {
    fmt.Println(Person{"Luiz Filipy", 19})

    p1 := Person{"Luiz", 19}
    p1.name = "Filipy"
    fmt.Println(p1)

    luiz := &p1
    fmt.Println(luiz.name)
}
