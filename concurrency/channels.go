package main

import (
    "fmt"
    "time"
    "sync"
)

/*Channel
serve para Goroutines se comunicam

data := <- a // read from channel a
a <- data // write to channel a

Os envios e recebimentos para um canal são bloqueando por padrão
*/

func test(done chan string) {
    fmt.Println("test")
    for i := 0; i < 10; i++ {
        fmt.Println("test ", i)
        time.Sleep(1 * time.Second)
    }
    done <- "ok"
}

func syncTest(done chan string, w *sync.WaitGroup) {
    fmt.Println("test")
    for i := 0; i < 10; i++ {
        fmt.Println("sync ", i)
        time.Sleep(1 * time.Second)
    }
    done <- "ok"
    w.Done()
}

func sendData(sendch chan<- int) {
    sendch <- 10
}

func server1(ch chan string) {
    time.Sleep(6 * time.Second)
    ch <- "from server1"
}
func server2(ch chan string) {
    time.Sleep(3 * time.Second)
    ch <- "from server2"

}

func main() {
    // Wait
    done := make(chan string)
    var wg sync.WaitGroup
    wg.Add(1)
    go syncTest(done, &wg)
    res3 := <-done
    fmt.Println(res3)

    // nao buferizado
    done1 := make(chan string)
    go test(done1)
    res1 := <-done1
    fmt.Println(res1)

    chnl1 := make(chan int)
    go sendData(chnl1)
    fmt.Println(<-chnl1)

    // buferizado
    done2 := make(chan string, 3)
    go test(done2)
    res2 := <-done2
    fmt.Println(res2)

    for i := 0; i < 10; i++ {
        fmt.Println("main ", i)
        time.Sleep(1 * time.Second)
    }

    output1 := make(chan string)
    output2 := make(chan string)
    go server1(output1)
    go server2(output2)
    select {
    case s1 := <-output1:
        fmt.Println(s1)
    case s2 := <-output2:
        fmt.Println(s2)
    default:
        fmt.Println("no value received")
    }
}
