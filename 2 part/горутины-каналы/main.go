// Канал позволяет отправлять и получать данные от горутин, а также обеспечивающая их взаимодействие и синхронизацию
// Не буферизированный канал - не имеет места хранения штук которые к нему приходят
// Буферизированный канал - имеет место хранения штук которые к нему приходят
package main

import (
	"fmt"
	"time"
)

func mine(transferPoint chan int, n int) {
	fmt.Println("ID for mine job: ", n, "started")
	time.Sleep(1 * time.Second)
	fmt.Println("ID for mine job: ", n, "is done")

	transferPoint <- 10
	fmt.Println("ID for mine job: ", n, "send to chanel")
}

func main() {
	coal := 0

	// transferPoint := make(chan int) // Не буферизированный канал
	transferPoint := make(chan int, 3) // Буферизированный канал

	initTime := time.Now()

	go mine(transferPoint, 1)
	go mine(transferPoint, 2)
	go mine(transferPoint, 3)

	coal += <-transferPoint
	coal += <-transferPoint
	coal += <-transferPoint

	fmt.Println("Count of coal: ", coal)
	fmt.Println("Time taken: ", time.Since(initTime))
}
