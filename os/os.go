package main

import (
	"fmt"
	"os"
	"os/user"
)

func main() {
	// retorna um nome de caminho raiz correspondente ao diretório atual.
	fmt.Println(os.Getwd())

	// retorna a identificação do processo do chamador.
	fmt.Println(os.Getpid())

	// retorna o nome do host
	fmt.Println(os.Hostname())

	// retorna o diretorio home do usuario atual
	fmt.Println(os.UserHomeDir())

	usr, _ := user.Current()
	fmt.Println(usr.Gid, usr.HomeDir, usr.Name, usr.Username, usr.Uid)

	fmt.Println(os.Args)
}
