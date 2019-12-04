package main

import (
	"fmt"
	"time"
)

func duration() {
	t0 := time.Now()
	fmt.Println("ok")
	t1 := time.Now()
	d := t1.Sub(t0)
	fmt.Printf("The call took %v to run.\n", d)
	fmt.Println(d.Milliseconds())
}

func TimeTest() {
	duration()

	t := time.Now()
	fmt.Printf("%d/%d/%d %d:%d:%d\n", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second())
}
