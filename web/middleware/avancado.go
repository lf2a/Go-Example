package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

// mostra um log no console
func Logging() Middleware {

	// cria um novo middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// define um http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// corpo do middleware
			start := time.Now()
			defer func() { log.Println(r.URL.Path, time.Since(start)) }()

			// chama o proximo middleware ou handler
			f(w, r)
		}
	}
}

// define se a url pode ser requisitada dado o tipo do metodo,
// se for diferente do valor passado irá retorna um 404
func Method(m string) Middleware {

	// cria um novo middleware
	return func(f http.HandlerFunc) http.HandlerFunc {

		// define um http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			// corpo do middleware
			if r.Method != m {
				http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
				return
			}

			// chama o proximo middleware ou handler
			f(w, r)
		}
	}
}

// aplica os middlewares para http.HandlerFunc
func Chain(f http.HandlerFunc, middlewares ...Middleware) http.HandlerFunc {
	for _, m := range middlewares {
		f = m(f)
	}
	return f
}

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello world")
}

func avancado() {
	http.HandleFunc("/", Chain(Hello, Method("GET"), Logging()))
	http.ListenAndServe(":8000", nil)
}
