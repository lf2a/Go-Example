package main

import "fmt"

func test()  {
    defer fmt.Println("World")

    fmt.Println("Hello")
}

func main()  {
    test()

    for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
