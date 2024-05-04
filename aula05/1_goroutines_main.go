package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
		time.Sleep(time.Second)
	}
}

func goRoutinesMain() {
	f("direct")
	go f("goroutine")

	time.Sleep(time.Second)
	fmt.Println("done")
}
