package main

import (
    "html/template"
    "net/http"
)

/*
   Control Structure               Definition
{{\* a comment *\/}} 	           Cria um comentario
{{.}} 	                           Renderiza o elemento raiz
{{.Title}}                         Renderiza o valor do campo `Title`
{{if .Done}} {{else}} {{end}} 	   Define uma declaração if-else
{{range .Todos}} {{.}} {{end}} 	   Define um loop
{{block "content" .}} {{end}} 	   Define um block com o nome `content`
*/

type Usuario struct {
    Nome  string
    Email string
    Idade byte
}

type Pagina struct {
    Titulo  string
    Usuario []Usuario
}

func verUsuarios(w http.ResponseWriter, r *http.Request) {
    // o arquivo a ser analisado para fazer o parse
    template := template.Must(template.ParseFiles("templates.html"))
    detalhes := []Usuario{
        {Nome: "Luiz Filipy", Email: "luiz@email.com", Idade: 20},
        {Nome: "Ana Paula", Email: "anapaula@email.com", Idade: 30},
        {Nome: "José Silva", Email: "jose_silva@email.com", Idade: 23},
    }

    data := Pagina{Titulo: "Golang", Usuario: detalhes}

    // irá retorna o html com as informações e colocando `Content-Type: text/html; charset=utf-8` no header
    template.Execute(w, data)
}

func templates() {
    http.HandleFunc("/", verUsuarios)
    http.ListenAndServe(":8000", nil)
}
