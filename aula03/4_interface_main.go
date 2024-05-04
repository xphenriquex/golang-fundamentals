package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Shape interface {
	Area() float64
}

func interfaceMain() {
	// circle := Circle{Radius: 10}
	// fmt.Println(circle.Area())

	rectangle := Rectangle{Height: 10, Width: 5}
	// fmt.Println(rectangle.Area())

	var shape Shape
	shape = rectangle
	fmt.Println(shape.Area())
}
