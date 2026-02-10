package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func job(ctx context.Context) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			fmt.Println("exit")
			return
		case <-time.After(1 * time.Second):
			fmt.Println("Application job:", i)
			i++
		}
	}
}

func main() {
	fmt.Println("PID:", os.Getpid())

	//sigchan := make(chan os.Signal, 1)
	//signal.Notify(sigchan, syscall.SIGTERM)
	//
	//sigchan1 := make(chan os.Signal, 1)
	//signal.Notify(sigchan1, syscall.SIGINT)
	//
	//go func() {
	//	for {
	//		s := <-sigchan
	//		fmt.Println("Got SIGTERM signal", s)
	//	}
	//}()
	//
	//go func() {
	//	for {
	//		s := <-sigchan1
	//		fmt.Println("Got SIGINT signal", s)
	//	}
	//}()

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGTERM)

	job(ctx)

	fmt.Println("Application stopped correctly")
}
