package main

import (
	"errors"
	"fmt"
	"log"
)

func soma(v1, v2 int32) (int32, error) {
	if v1 == 0 && v2 == 0 {
		return -1, errors.New("Os valores devem ser diferentes de zero")
	}

	return v1 + v2, nil
}

func erro() {
	res, err := soma(12, 2)

	if err != nil {
		log.Panicln(err)
	}

	fmt.Println("res", res)

	res2, err2 := soma(0, 0)

	if err2 != nil {
		log.Panicln(err2)
	}

	fmt.Println("res2:", res2)
}
