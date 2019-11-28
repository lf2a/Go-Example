package main

import "fmt"

func add(v1, v2 int) int {
    return v1 + v2
}

// retorna multiplos valores
func name(age int, name string ) (int, string) {
    return age, name
}

// retorna valores que podem ser nomeados
func person(n string, a int) (name string, age int) {
    name = "is " + n
    age = a
    return
}

// function arg
func calc(fn func(v1, v2 int) int) int {
    return fn(3, 4)
}

func test() func(float64) float64 {
    return func(y float64) float64 {
        return y / 2
    }
}

func variadic(num int, nums ...int)  {
    fmt.Println(num, nums)
}

func main()  {
    fmt.Println(add(23, 2))

    a, n := name(19, "Luiz")
    fmt.Println(a, n)

    fmt.Println(person("Luiz", 19))

    sum := func(v1, v2 int) int {
        return v1 + v2
    }
    fmt.Println(calc(sum))
    fmt.Println(sum(5, 9))

    t1 := test()
    fmt.Println(t1(12))

    variadic(12, 32, 65, 56, 7, 89)
}
