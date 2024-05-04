package main

import "fmt"

func deferMain() {
	defer fmt.Println("World")
	fmt.Println("Hello")
}
