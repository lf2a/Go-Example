package main

import "fmt"

func main()  {
    age := 18
    switch age {
    case 15:
        fmt.Println("Menor de idade")
    case 18:
        fmt.Println("Maior de idade")
    default:
        fmt.Println("Inválido")
    }


    age = 17
    switch  {
    case age < 18:
        fmt.Println("Menor de idade")
    case age >= 18:
        fmt.Println("Maior de idade")
    default:
        fmt.Println("Inválido")
    }
}
