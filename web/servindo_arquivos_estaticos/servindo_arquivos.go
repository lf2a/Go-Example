package main

import (
    // "fmt"
	"net/http"
)

/*
	fs := http.FileServer(http.Dir("static/"))
	Serve para veicular ativos estáticos como JavaScript, CSS e imagens,
	usamos o http.FileServer embutido e o apontamos para um caminho de URL.
	Para que o servidor de arquivos funcione corretamente,
	é necessário saber de onde servir os arquivos.

	http.Handle("/static/", http.StripPrefix("/static/", fs))
	Uma vez que nosso servidor de arquivos esteja instalado,
	basta apontar um caminho de URL para ele, assim como fizemos
	com as solicitações dinâmicas. Uma coisa a ser observada: para
	servir os arquivos corretamente, precisamos remover uma parte do
	caminho da URL. Normalmente, esse é o nome do diretório em que
	nossos arquivos ficam.
*/

func servindoArquivos() {
    fs := http.FileServer(http.Dir("./static"))
    http.Handle("/", http.StripPrefix("/", fs))

    http.ListenAndServe(":8000", nil)
}
