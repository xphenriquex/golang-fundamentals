package main

import "fmt"

func anonymousMain() {
	area := func(width, height float32) float64 {
		return float64(width * height)
	}
	fmt.Println(area(2, 3))

	areResult := func(width, height float64) float64 {
		return width * height
	}(2, 3)
	fmt.Println(areResult)
}
