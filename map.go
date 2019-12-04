package main

import "fmt"

type person2 struct {
	age uint8
}

func MapTest() {
	var m map[string]person2

	m = make(map[string]person2)
	m["luiz"] = person2{19}
	m["filipy"] = person2{30}

	fmt.Println(m["luiz"])
	fmt.Println(m)

	delete(m, "filipy")
	fmt.Println(m)

	mt, ok := m["filipy"]
	fmt.Println(mt, ok)

	var mm = map[string]person2{
		"luiz":   person2{56},
		"filipy": person2{78},
	}
	// or
	// var mm = map[string]person2{
	//     "luiz": {56},
	//     "filipy": {78},
	// }

	fmt.Println(mm)
}
