package main

import (
	"context"
	"example/miner"
	"example/postman"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var coal atomic.Int64

	mtx := sync.Mutex{}
	var mails []string

	minerContext, minerCancel := context.WithCancel(context.Background())
	postmanContext, postmanCancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("---Miner is done!---")
		minerCancel()
	}()
	go func() {
		time.Sleep(6 * time.Second)
		fmt.Println("---Postman is done!---")
		postmanCancel()
	}()

	coalTransferPoint := miner.MinerPool(minerContext, 100)
	mailTransferPoint := postman.PostmanPool(postmanContext, 100)

	initTime := time.Now()

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for v := range coalTransferPoint {
			coal.Add(int64(v))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for v := range mailTransferPoint {
			mtx.Lock()
			mails = append(mails, v)
			mtx.Unlock()
		}
	}()

	wg.Wait()

	//isCoalClosed := false
	//isMailClosed := false
	//for !isCoalClosed || !isMailClosed {
	//	select {
	//	case c, ok := <-coalTransferPoint:
	//		if !ok {
	//			isCoalClosed = true
	//			continue
	//		}
	//		coal += c
	//	case m, ok := <-mailTransferPoint:
	//		if !ok {
	//			isMailClosed = true
	//			continue
	//		}
	//		mails = append(mails, m)
	//	}
	//}

	fmt.Println("coal:", coal.Load())

	mtx.Lock()
	fmt.Println("mails:", len(mails))
	mtx.Unlock()

	fmt.Println("Time elapsed: ", time.Since(initTime))
}
