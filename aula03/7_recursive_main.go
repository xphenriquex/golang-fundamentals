package main

import (
	"fmt"
)

func printDecreasing(number int) {
	fmt.Println(number)
	if number > 0 {
		printDecreasing(number - 1)
	}
}

func recursiveMain() {
	printDecreasing(10)
}
