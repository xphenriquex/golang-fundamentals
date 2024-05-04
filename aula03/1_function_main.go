package main

import "fmt"

func rectangleArea(width, height float64) (result float64, result2 float64) {
	result = width * height
	result = 1.1
	return
}

func functionMain() {
	res1, _ := rectangleArea(2, 3)
	fmt.Println("rectangleArea =", res1)
	// fmt.Println("rectangleArea = ", rectangleArea(2, 3))
}
