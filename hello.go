package main

import "fmt"
import "os"
import "net/http"

func main() {

	boasVindas()
	exibeMenu()

	comando := leComando()

	switch comando {
	case 1:
		iniciarMonitoramento()
	case 2:
		fmt.Println("Exibindo os logs")
	case 0:
		fmt.Println("Saindo")
		os.Exit(0)
	default:
		fmt.Println("Opção inválida!")
		os.Exit(-1)
	}

}

func devolveNomeCompleto() (string, string) {
	nome := "Elvis"
	sobreNome := "Estevan"
	return nome, sobreNome
}

func exibeMenu() {
	fmt.Println("\n1. Iniciar monitoramento")
	fmt.Println("2. Exibir os logs")
	fmt.Println("0. Sair")
	fmt.Print("\nInforme: ")
}

func leComando() int {
	var comando int
	fmt.Scan(&comando)

	fmt.Println(comando)
	return comando
}

func boasVindas() {
	nome, sobreNome := devolveNomeCompleto()
	versao := 1.1

	fmt.Println("Olá, sr.", nome, sobreNome)
	fmt.Println("Este programa está na versão", versao)
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando")
	site := "https://alura.com.br"
	response, _ := http.Get(site)
	fmt.Println(response)
}