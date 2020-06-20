package main

/*
	O Cgo permite que os pacotes Go chamem código C.
	Dado um arquivo de origem Go escrito com alguns recursos especiais,
	o cgo gera arquivos Go e C que podem ser combinados em um único pacote Go.
*/

/*
#include <stdio.h>
#include <stdlib.h>

static void myprint(char* s) {
  printf("%s\n", s);
}

int Poww(int x, int y) {
	int temp = 1;
	for(int i = 1; i <= y; i++){
		printf("i:%d x:%d y:%d", i, x, y);
		temp = x * temp;
	}

	return temp;
}

float Squaree(float x, float y) {
	return x / y;
}
*/
import "C"
import (
	"fmt"
	"reflect"
	"unsafe"
)

/*
=======================================================
C 						 - CGO             - Go
=====================================================
char  					 - C.char          - byte
-----------------------------------------------------
singed char 			 - C.schar         - int8
-----------------------------------------------------
unsigned char			 - C.uchar         - uint8
-----------------------------------------------------
short 					  - C.short        - int16
-----------------------------------------------------
unsigned short 			  - C.ushort       - uint16
-----------------------------------------------------
int 					  - C.int          - int32
-----------------------------------------------------
unsigned int 			  - C.uint         - uint32
-----------------------------------------------------
long 					  - C.long         - int32
-----------------------------------------------------
unsigned long 			  - C.ulong        - uint32
-----------------------------------------------------
long long int 			  - C.longlong 	   - int64
-----------------------------------------------------
unsigned long long int    - C.ulonglong    - uint64
-----------------------------------------------------
float 					  - C.float 	   - float32
-----------------------------------------------------
double 					  - C.double 	   - float64
-----------------------------------------------------
size_t 					  - C.size_t 	   - uint
-----------------------------------------------------
*/

func cgo_exemplo() {
	// executando exemplo 1

	// A conversão entre strings de Go e C é feita com as funções C.CString, C.GoString e C.GoStringN.
	// Essas conversões fazem uma cópia dos dados da string.
	// `nome` é um ponteiro
	nome := C.CString("Luiz Filipy")

	// chama a função `myprint(char* s)` e mostrar o conteudo da variável `nome`
	C.myprint(nome)

	// mostra o tipo de dados da variavel `nome`
	fmt.Println(reflect.TypeOf(nome), "\n\n")

	// A chamada para C.CString retorna um ponteiro para o início do array de caracteres, portanto,
	// antes da função ser encerrada, a converteremos em um `unsafe.Pointer()` e liberaremos a
	// alocação de memória com `C.free`.
	C.free(unsafe.Pointer(nome))

	// excutando exemplo 2
	// chama a função `Poww(int x, int y)` e converte o tipo int de go para o tipo int de c
	res := C.Poww(C.int(2), C.int(3))
	fmt.Println(res)

	// mostra o tipo retornado
	fmt.Println(reflect.TypeOf(res), "\n\n")

	// executando exemplo 3
	// chama a função `Poww(int x, int y)` e converte o tipo float de go para o tipo float de c
	res2 := C.Squaree(C.float(34.7), C.float(23.8))
	fmt.Println(res2)

	// mostra o tipo retornado
	fmt.Println(reflect.TypeOf(res2))
}
