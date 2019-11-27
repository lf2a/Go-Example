package main

import "fmt"

type Person struct {
    age uint8
}

func main() {
    var m map[string]Person

    m = make(map[string]Person)
    m["luiz"] = Person{19}
    m["filipy"] = Person{30}

    fmt.Println(m["luiz"])
    fmt.Println(m)

    delete(m, "filipy")
    fmt.Println(m)

    mt, ok := m["filipy"]
    fmt.Println(mt, ok)


    var mm = map[string]Person{
        "luiz": Person{56},
        "filipy": Person{78},
    }
    // or
    // var mm = map[string]Person{
    //     "luiz": {56},
    //     "filipy": {78},
    // }

    fmt.Println(mm)
}
