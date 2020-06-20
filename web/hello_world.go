package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/*
   func (w http.ResponseWriter, r *http.Request)
   Manipulador que recebe todas as conexões HTTP recebidas
   de navegadores, clientes HTTP ou solicitações de API.

   A função recebe dois parâmetros:

   Um http.ResponseWriter no qual você escreve sua resposta de texto/html.
   Um http.Request que contém todas as informações sobre esta solicitação HTTP,
   incluindo itens como os campos URL ou cabeçalho.

   http.ListenAndServe(":8000", nil) é usado para o seridor poder
   escutar todas as requisicoes na porta 8000
*/

type Pessoa struct {
	Nome string
}

func helloUsuario(w http.ResponseWriter, r *http.Request) {
	nome := r.URL.Query()["nome"][0]
	p := Pessoa{
		Nome: nome,
	}
	// retornando um json
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func hello() {
	// HandleFunc(): recebe dois parametros, a rota e a função de disparo dessa rota
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World!")
		fmt.Println(r)
	})
	http.HandleFunc("/usuario", helloUsuario)

	http.ListenAndServe(":8000", nil)
}
