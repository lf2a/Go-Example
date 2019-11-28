package main

import (
    "fmt"
    "time"
)

func test()  {
    for i := 0;i < 10;i++ {
        fmt.Println(i)
        time.Sleep(1000 * time.Millisecond)
    }
}

func main()  {
    go test()
    go test()

    test()
}
