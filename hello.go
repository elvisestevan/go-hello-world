package main

import "fmt"

func main() {
	nome := "Elvis"
	versao := 1.1

	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("\n1. Iniciar monitoramento")
	fmt.Println("2. Exibir os logs")
	fmt.Println("0. Sair")
	fmt.Print("\nInforme: ")

	var comando int
	fmt.Scan(&comando)

	fmt.Println(comando)

}
