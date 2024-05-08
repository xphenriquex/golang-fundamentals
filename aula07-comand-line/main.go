package main

import (
	"fmt"
	"log"
	"os"

	"github.com/xphenriquex/golang-aulas/aula07-comand-line/app"
)

func main() {
	fmt.Println("Ponto de partida")

	aplicacao := app.Gerar()
	if erro := aplicacao.Run(os.Args); erro != nil {
		log.Fatal(erro)
	}
}
