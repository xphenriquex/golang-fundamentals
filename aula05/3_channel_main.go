package main

import (
	"fmt"
	"time"
)

func sum(values []int, c chan int) {
	sum := 0
	for _, v := range values {
		sum += v
	}
	c <- sum // send sum to channel c
}

func channelMain() {
	// numbers := new([]int) //ponteiros
	// numbers := make([]int, 0)
	// fmt.Println(numbers)

	// messages := make(chan string)
	// var value int
	// go func(m chan string, v *int) {
	// 	time.Sleep(3 * time.Second)
	// 	m <- "buffered"
	// 	*v = 10
	// }(messages, &value)
	// fmt.Println(<-messages, value)

	//ch <- value //Inserir
	//<-ch 	      //obter

	//------------------------------------------------------------
	// values := []int{10, 20, 30, -20, 0, 10}

	// c := make(chan int)
	// go sum(values[len(values)/2:], c)
	// go sum(values[:len(values)/2], c)
	// x, y := <-c, <-c //receive from channel c

	// fmt.Println(x, y, x+y)

	//------------------------------------------------------------

	channel := make(chan string)
	go writeOnscreen("Olá Mundo!", channel)

	fmt.Println("Depois da função escrever começar a ser executada")

	// for {
	// 	message, open := <-channel
	// 	if !open {
	// 		break
	// 	}
	// 	fmt.Println(message)
	// }

	for message := range channel {
		fmt.Println(message)
	}

	fmt.Println("Fim do programa")
}

func writeOnscreen(text string, channel chan string) {
	for i := 0; i < 5; i++ {
		channel <- text
		time.Sleep(time.Second)
	}
	close(channel)
}
