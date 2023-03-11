package saudacao

import "fmt"

func Ola_a(nome string) string {
	mensagem := fmt.Sprintf("Olá, %v. Seja bem vindo!", nome)
	return mensagem
}
