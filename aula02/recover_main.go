package main

import (
	"fmt"
	"log"
)

func RecoverMain() {
	divideByZero()
	fmt.Println("we survived diving by zero!")
}

func divideByZero() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("panic occurred:", err)
		}
	}()
	fmt.Println(divide(1, 0))
}

func divide(a, b int) int {
	return a / b
}
