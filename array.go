package main

import "fmt"

func getInfo(s []string) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func ArrayTest() {
	var a [5]int
	a[0] = 19
	a[1] = 34
	fmt.Println(a)

	// nil
	var x []int
	fmt.Println(x, len(x), cap(x))
	if x == nil {
		fmt.Println("nil!")
	}

	countries := []string{"Brasil", "EUA", "Canada", "Mexico"}
	fmt.Println(countries)
	getInfo(countries[:2])
	getInfo(countries)

	// slice
	countries2 := countries[0:2]
	fmt.Println(countries2)

	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)

	// struct array
	persons := []struct {
		name string
		age  uint8
	}{
		{"Luiz", 45},
		{"Carla", 19},
		{"Pedro", 23},
		{"Lua", 64},
		{"José", 34},
		{"Maria", 12},
	}
	fmt.Println(persons)

	// make(type, 0)
	m := make([]int, 5)
	m[0] = 12
	fmt.Println(m, len(m), cap(m))

	// append
	var z []int
	fmt.Println(z, len(z), cap(z))

	z = append(z, 1, 2)
	fmt.Println(z, len(z), cap(z))

	z = append(z, 3, 4)
	fmt.Println(z, len(z), cap(z))

	z = append(z, 5, 6)
	fmt.Println(z, len(z), cap(z))
}
