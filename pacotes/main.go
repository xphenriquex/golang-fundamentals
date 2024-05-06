package main

import (
	"fmt"

	"github.com/badoux/checkmail"
	"github.com/xphenriquex/golang-aulas/pacotes/auxiliar"
)

func main() {
	fmt.Println("Escrevendo do arquivo man")
	auxiliar.Escrever()

	// erro := checkmail.ValidateFormat("devbook@gmail.com")
	erro := checkmail.ValidateFormat("123")
	fmt.Println(erro)
}
