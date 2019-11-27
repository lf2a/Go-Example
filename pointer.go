package main

import "fmt"

var a, b int = 12, 19

func main() {
    fmt.Println(a, b)

    aa := &a
    bb := &b
    fmt.Printf("%d %d\n",aa, bb)
    fmt.Printf("%d %d\n", *aa, *bb)

    *aa, *bb = 45, 78
    fmt.Println(a, b)
}
