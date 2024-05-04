package main

import (
	"fmt"
	"sync"
	"time"
)

func f2(from string, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(from, ": starting")
	time.Sleep(time.Second)
	fmt.Println(from, ": done")
}

func waitGroupsMain() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go f2(fmt.Sprintf("goroutine %d", i), &wg)
	}

	wg.Wait()
	fmt.Println("done")
}
