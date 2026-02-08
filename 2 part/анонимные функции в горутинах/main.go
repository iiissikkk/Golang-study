package main

import (
	"fmt"
	"time"
)

func foo() {
	for {
		fmt.Println("Foo")
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go foo()

	go func() {
		for {
			fmt.Println("Anon")
			time.Sleep(750 * time.Millisecond)
		}
	}()

	time.Sleep(5 * time.Second)
}
