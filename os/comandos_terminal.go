package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	// Command: retorna a estrutura Cmd para executar o programa nomeado com os argumentos fornecidos
	cmd := exec.Command("ls", "-l")

	var out bytes.Buffer
	cmd.Stdout = &out

	// Run: inicia o comando especificado e aguarda a conclusão.
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(out.String())

	// ========= exemplo 2

	// Output: executa o comando e retorna sua saída padrão.
	out2, err2 := exec.Command("date").Output()
	if err2 != nil {
		log.Fatal(err2)
	}

	fmt.Printf("A data de hoje é: %s\n", out2)

	// ========== exemplo 3

	cmd3 := exec.Command("sleep", "1")

	fmt.Println("Esperando...")

	err3 := cmd3.Run()
	if err3 != nil {
		log.Println(err3)
	}

	// ========= exemplo 4

	cmd4 := exec.Command("sleep", "5")

	// Start inicia o comando especificado, mas não espera que seja concluído.
	err4 := cmd4.Start()
	if err4 != nil {
		log.Fatal(err4)
	}

	fmt.Printf("Esperando o comando terminar...")

	// Wait: aguarda o comando sair e aguarda a cópia de stdin ou a cópia de stdout ou stderr para ser concluída.
	err4 = cmd4.Wait()
	fmt.Printf("Erros: %v", err4)
}
