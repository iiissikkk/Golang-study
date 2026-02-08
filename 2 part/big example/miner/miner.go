package miner

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Miner(ctx context.Context, wg *sync.WaitGroup, transferPoint chan<- int, n int, power int) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("I am a Miner:", n, "finished work")
			return
		default:
			fmt.Println("I am a Miner:", n, "start working")
			time.Sleep(1 * time.Second)
			fmt.Println("I am a Miner:", n, "Coal is mined")

			transferPoint <- power
			fmt.Println("I am a Miner:", n, "Coal is send", power)
		}
	}
}

func MinerPool(ctx context.Context, minerCount int) <-chan int {
	coalTransferPoint := make(chan int)

	wg := &sync.WaitGroup{}

	for i := 0; i <= minerCount; i++ {
		wg.Add(1)
		go Miner(ctx, wg, coalTransferPoint, i, rand.Intn(100))
	}

	go func() {
		wg.Wait()
		close(coalTransferPoint)
	}()

	return coalTransferPoint
}
