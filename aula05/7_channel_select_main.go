package main

import (
	"fmt"
	"time"
)

func fibonacci2(channel, exit chan int) {
	x, y := 0, 1
	for {
		select {
		case channel <- x:
			x, y = y, x+y
		case <-exit:
			fmt.Println("exit")
			return
		}
	}
}

func channelSelectMain() {
	// channel := make(chan int)
	// exit := make(chan int)
	// go func() {
	// 	for i := 0; i < 8; i++ {
	// 		fmt.Println(<-channel)
	// 	}
	// 	exit <- 0
	// }()

	// fibonacci2(channel, exit)

	//----------------------------------------------------

	channel1, channel2 := make(chan string), make(chan string)

	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			channel1 <- "Canal 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			channel2 <- "Canal 2"
		}
	}()

	for {
		select {
		case messageChannel1 := <-channel1:
			fmt.Println(messageChannel1)
		case messageChannel2 := <-channel2:
			fmt.Println(messageChannel2)
		}
	}
}
