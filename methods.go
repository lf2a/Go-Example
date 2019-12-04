package main

import "fmt"

type Person struct {
	name string
	age  uint8
}

func (p *Person) getInfo() {
	fmt.Printf("%s tem %d anos\n", p.name, p.age)
}

func getInfo2(p *Person) {
	fmt.Printf("%s tem %d anos\n", p.name, p.age)
}

type myBigInt int64

func (b *myBigInt) test() {
	fmt.Println(*b, b)
}

func MethodTest() {
	p1 := &Person{"Luiz", 19}
	p1.getInfo()

	getInfo2(p1)

	b := myBigInt(1234567890)
	b.test()
}
