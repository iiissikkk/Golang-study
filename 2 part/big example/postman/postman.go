package postman

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Postman(ctx context.Context, wg *sync.WaitGroup, transferPoint chan<- string, n int, mail string) {
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Println("I am a Postman:", n, "finished work")
			return
		default:
			fmt.Println("I am a Postman:", n, "start working")
			time.Sleep(1 * time.Second)
			fmt.Println("I am a Postman:", n, "Mail is send")

			transferPoint <- mail
			fmt.Println("I am a Postman:", n, "Mail is given", mail)
		}
	}
}

func PostmanPool(ctx context.Context, postmanCount int) <-chan string {
	mailTransferPoint := make(chan string)

	wg := &sync.WaitGroup{}

	for i := 0; i <= postmanCount; i++ {
		wg.Add(1)
		go Postman(ctx, wg, mailTransferPoint, i, postmanToMail(i))
	}

	go func() {
		wg.Wait()
		close(mailTransferPoint)
	}()

	return mailTransferPoint
}

func postmanToMail(postmanNumber int) string {
	ptm := map[int]string{
		1: "Family business",
		2: "Friend's hb",
		3: "Some info",
	}

	mail, ok := ptm[postmanNumber]
	if !ok {
		return "Lotery!"
	}

	return mail
}
