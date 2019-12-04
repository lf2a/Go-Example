package main

import (
	"fmt"
	"math/cmplx"
)

/*Types
- bool
- string
- int  int8  int16  int32  int64
- uint uint8 uint16 uint32 uint64 uintptr
- byte -> alias for `uint8`
- rune -> alias for `int32` represents a Unicode code point
- float32 float64
- complex64 complex128
*/

// The expression T(v) converts the value v to the type T.

var aA, bB, cC bool
var day, month, year int = 27, 11, 2019
var (
	toBe          = false
	maxInt uint64 = 1<<64 - 1
	z             = cmplx.Sqrt(-5 + 12i)
)

const pi = 3.14 // const variable

func VariableTest() {
	var len int
	fmt.Println(aA, bB, cC, len)

	fmt.Printf("%d/%d/%d\n", day, month, year)

	// := só pode ser dentro de funções
	i, j := 12, 20
	fmt.Println(i, j)

	fmt.Printf("Type: %T Value: %v\n", toBe, toBe)
	fmt.Printf("Type: %T Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
