package main

import "fmt"

type iPerson interface {
	getInfo()
}

type person1 struct {
	name string
	age  uint8
}

func Interface() {
	var ip iPerson
	p1 := person1{"luiz", 19}

	ip = &p1
	ip.getInfo()
	var p2 iPerson = person1{"luiz", 19}
	p2.getInfo()
}

func (p person1) getInfo() {
	fmt.Println(p.name, p.age)
}
