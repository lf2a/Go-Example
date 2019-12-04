package main

import (
	"errors"
	"fmt"
	"log"
)

func add(v1, v2 int32) (int32, error) {
	if v1 == 0 && v2 == 0 {
		return -1, errors.New("Os valores devem ser diferentes de zero")
	}
	return v1 + v2, nil
}

func ErrorTest() {
	res1, err1 := add(12, 2)
	if err1 != nil {
		log.Panicln(err1)
		// fmt.Println(err1)
	} else {
		fmt.Println(res1)
	}

	res, err := add(0, 0)
	if err != nil {
		log.Panicln(err)
		// fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
