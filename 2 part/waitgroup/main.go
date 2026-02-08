package main

import (
	"fmt"
	"sync"
	"time"
)

func postman(wg *sync.WaitGroup, text string) {
	defer wg.Done()

	for i := 0; i <= 3; i++ {
		fmt.Println("Send mail", text, "for the", i, "time")
		time.Sleep(250 * time.Millisecond)
	}
}

func main() {
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go postman(wg, "News")

	wg.Add(1)
	go postman(wg, "Playboy")

	wg.Add(1)
	go postman(wg, "Golang")

	wg.Wait()

	fmt.Println("done")
}
