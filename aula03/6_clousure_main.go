package main

import (
	"fmt"
)

func sequency() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func clousureMain() {
	nexInt := sequency()
	fmt.Println(" newInt =", nexInt())
	fmt.Println(" newInt =", nexInt())
	fmt.Println(" newInt =", nexInt())

	newInteger := sequency()
	fmt.Println(" newInteger =", newInteger())

	fmt.Println(" newInt =", nexInt())

	fmt.Println(" sequency =", sequency()())
	fmt.Println(" sequency =", sequency()())
	fmt.Println(" sequency =", sequency()())
}
