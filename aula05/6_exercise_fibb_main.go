package main

import (
	"fmt"
)

// func sumNumber(ch chan int) {
// 	sum := 0
// 	temp := 0
// 	before := 1
// 	after := 1
// 	for i := 0; i < 8; i++ {
// 		if i <= 1 {
// 			sum = i
// 		} else {
// 			if i < 4 {
// 				after += temp
// 				temp++
// 				sum = after
// 			} else {
// 				sum = before + after
// 				before = after
// 				after = sum
// 			}
// 		}

// 		ch <- sum
// 	}
// 	close(ch)
// }

func fibonacci(ch chan int, number int) {
	arr := []int{0, 1, 1, 2}
	sum := 0
	before, after := 1, 2
	for i := 0; i < number; i++ {
		if i < 4 {
			sum = arr[i]
		} else {
			sum = before + after
			before = after
			after = sum
		}
		ch <- sum
	}
	close(ch)
}

func exerciseMain() {
	channel := make(chan int, 1)
	number := 8
	go fibonacci(channel, number)

	for v := range channel {
		fmt.Print(v, " ")
	}
}

// func fibonacci(number int, channel chan int) {
// 	x, y := 0, 1
// 	for i := 0; i < number; i++ {
// 		channel <- x
// 		x, y = y, x+y
// 	}
// 	close(channel)
// }

// func exerciseMain() {
// 	channel := make(chan int, 8)
// 	go fibonacci(cap(channel), channel)

// 	for i := range channel {
// 		fmt.Println(i)
// 	}
// }
