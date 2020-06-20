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
	Idade string
}

type Pagina struct {
	Titulo  string
	Usuario Usuario
}

func verUsuarios(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("form.html"))
	if r.Method != http.MethodPost {
		template.Execute(w, nil)
		return
	}
	datalhes := Usuario{
		Nome:  r.FormValue("nome"),
		Email: r.FormValue("email"),
		Idade: r.FormValue("idade"),
	}

	data := Pagina{Titulo: "Obrigado!", Usuario: datalhes}

	template.Execute(w, data)
}

func form() {
	http.HandleFunc("/", verUsuarios)
	http.ListenAndServe(":8000", nil)
}
