package main

import (
    "fmt"
    "net/http"
)

/* func (w http.ResponseWriter, r *http.Request)
Manipulador que recebe todas as conexões HTTP recebidas
de navegadores, clientes HTTP ou solicitações de API.

A função recebe dois parâmetros:

Um http.ResponseWriter no qual você escreve sua resposta de texto / html.
Um http.Request que contém todas as informações sobre esta solicitação HTTP,
incluindo itens como os campos URL ou cabeçalho.

http.ListenAndServe(":80", nil) é usado para o seridor poder
escutar todas as requisicoes na porta 80
*/

func main() {
    http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request)  {
        fmt.Fprintf(w, "Hello World!")
        fmt.Println(r)
    })

    http.ListenAndServe(":8000", nil)
}
