package main

import (
	"fmt"
	"time"
)

func write(ch chan int) {
	for i := 0; i < 4; i++ {
		ch <- i
		fmt.Println("successfully wrote", i, "to channel")
	}
	close(ch)
}

func channelBufferMain() {
	// messages := make(chan string, 2)
	// messages <- "buffered"
	// messages <- "channel"
	// fmt.Println(<-messages)
	// messages <- "other"
	// fmt.Println(<-messages)
	// fmt.Println(<-messages)

	// ----------------------------------------------------

	chanel := make(chan int, 2)
	go write(chanel)
	time.Sleep(2 * time.Second)
	for v := range chanel {
		fmt.Println("read value", v, "from channel")
		time.Sleep(2 * time.Second)
	}

	// Canais com buffer são úteis quando você sabe quantas goroutines você lançou,
	// quer limitar o número de goroutines que você vai lançar ou quer limitar a
	// quantidade de trabalho que está na fila.
}
