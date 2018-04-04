package main

import "fmt"
import "reflect"

func main() {
	nome := "Elvis"
	idade := 28
	versao := 1.1
	fmt.Println("Olá, sr.", nome, "sua idade é", idade, "anos")
	fmt.Println("Este programa está na versão", versao)

	fmt.Println("Tipo nome:", reflect.TypeOf(nome))
	fmt.Println("Tipo idade:", reflect.TypeOf(idade))
	fmt.Println("Tipo versao:", reflect.TypeOf(versao))
}
