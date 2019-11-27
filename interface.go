package main

import "fmt"

type IPerson interface {
    getInfo()
}

type Person struct {
    name string
    age uint8
}

func main() {
    // var ip IPerson
    // p1 := Person{"luiz", 19}
    //
    // ip = &p1
    // ip.getInfo()
    var p1 IPerson = Person{"luiz", 19}
    p1.getInfo()
}

func (p Person) getInfo()  {
    fmt.Println(p.name, p.age)
}
