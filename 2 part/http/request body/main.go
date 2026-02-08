package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"sync"
	"sync/atomic"
)

var mtx = sync.Mutex{}
var money = atomic.Int64{}
var bank = atomic.Int64{}

func payHandler(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Something went wrong! Body is not read", err)
		return
	}

	httpRequestBodyString := string(httpRequestBody)
	paymentAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		fmt.Println("Something went wrong! Fail to convert to int:", err)
		return
	}

	mtx.Lock()
	if money.Load()-int64(paymentAmount) >= 0 {
		money.Add(int64(-paymentAmount))
		fmt.Println("Payment was successes!", money.Load())
	} else {
		fmt.Println("Payment was fails!", money.Load())
	}
	mtx.Unlock()
}

func saveHandler1(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Something went wrong! Body is not read", err)
		return
	}

	httpRequestBodyString := string(httpRequestBody)
	saveAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		fmt.Println("Something went wrong! Fail to convert to int", err)
		return
	}

	mtx.Lock()
	if money.Load() >= int64(saveAmount) {
		money.Add(int64(-saveAmount))
		bank.Add(int64(saveAmount))
		fmt.Println("Save money:", money.Load())
		fmt.Println("Save amount in bank:", bank.Load())
	} else {
		fmt.Println("Nothing to save!", money.Load())
	}
	mtx.Unlock()
}

func main() {
	money.Add(1000)

	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/save", saveHandler1)

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Something went wrong! Server is not started", err)
	}
}
