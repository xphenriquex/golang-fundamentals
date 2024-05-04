package main

import (
	"fmt"
)

const (
	ano  = "ano"
	dois = "dois"
)

func main() {
	// var name string = "ph"
	// name = "ph" só atribuição
	// name := "ph" // criação e atribuição
	// fmt.Println(name)

	//ponteiros***********

	// var pname *string
	// s := "ssss"

	// pname = &s

	// fmt.Println(*pname)

	//map
	// mapa := map[string]string{}

	// mapa = nil

	// fmt.Println(mapa)

	// //slice
	// lista := []string{}
	// fmt.Println(lista)

	//printf
	// name := "name"

	// fmt.Printf("%s", name)

	//if

	// condicao := false

	// if condicao {
	// 	fmt.Println("ok")
	// } else {
	// 	fmt.Println("not ok")
	// }

	//for

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	//while
	// for {
	// 	fmt.Println("loop")
	// }

	//do while
	// condition := true
	// for condition {
	// 	fmt.Println("loop")
	// 	condition = false
	// }

	// for i, v := range []int{1, 2, 3} {
	// 	fmt.Println(i, v)
	// }

	// fat(0) = 1
	// fat(n) = 1 * 2 * ... * (n - 1)

	//switch
	// expression := 3
	// switch expression {
	// case 1:
	// 	fmt.Println(1)
	// case 2:
	// 	fmt.Println(2)
	// case 10:
	// 	fmt.Println(10)
	// default:
	// 	fmt.Println("nenhuma operação selecionada")
	// }

	// t := time.Now()
	// switch {
	// case t.Hour() < 12:
	// 	fmt.Println("manhã")
	// case t.Hour() < 18:
	// 	fmt.Println("tarde")
	// default:
	// 	fmt.Println("noite")
	// }

	// fatorial
	fat := 5
	result := fat

	for i := fat; i > 1; i-- {
		result = result * (i - 1)
	}

	if fat == 0 {
		fmt.Println(1)
	} else {
		fmt.Println(result)
	}

}
