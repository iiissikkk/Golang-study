// Оператор select используется для ожидания нескольких операций канала

package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	intCh := make(chan int)
	strCh := make(chan string)

	go func() {
		i := 1
		for {
			intCh <- i
			i++

			time.Sleep(200 * time.Millisecond)
		}
	}()

	go func() {
		i := 1
		for {
			strCh <- "hi " + strconv.Itoa(i)
			i++

			time.Sleep(200 * time.Millisecond)
		}
	}()

	for {
		select {
		case number := <-intCh:
			fmt.Println("IntCh:", number)
		case str := <-strCh:
			fmt.Println("StrCh:", str)
		}
	}
}
