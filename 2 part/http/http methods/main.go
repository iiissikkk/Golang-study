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
	// выводим http method
	fmt.Println("HTTP method:", r.Method)
	if r.Method != http.MethodPost {
		msg := "Method not allowed. Use POST"
		fmt.Println(msg)

		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte(msg))
		return
	}

	fmt.Println("OK: correct method POST")

	// выводим header
	for k, v := range r.Header {
		fmt.Println(k, v)
	}

	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		msg := "Something went wrong! Body is not read" + err.Error()
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error writing response:", err)
		}
		return
	}

	httpRequestBodyString := string(httpRequestBody)
	paymentAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		msg := "Something went wrong! Fail to convert to int:" + err.Error()
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error writing response:", err)
		}
		return
	}

	mtx.Lock()
	if money.Load()-int64(paymentAmount) >= 0 {
		money.Add(int64(-paymentAmount))
		msg := fmt.Sprintf("Payment was success! %d", money.Load())
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error writing response:", err)
		}
	} else {
		fmt.Println("Payment was fails!", money.Load())
	}
	mtx.Unlock()
}

func saveHandler1(w http.ResponseWriter, r *http.Request) {
	httpRequestBody, err := io.ReadAll(r.Body)
	if err != nil {
		msg := "Something went wrong! Body is not read" + err.Error()
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error writing response:", err)
		}
		return
	}

	httpRequestBodyString := string(httpRequestBody)
	saveAmount, err := strconv.Atoi(httpRequestBodyString)
	if err != nil {
		msg := "Something went wrong! Fail to convert to int:" + err.Error()
		fmt.Println(msg)
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error writing response:", err)
		}
		return
	}

	mtx.Lock()
	if money.Load() >= int64(saveAmount) {
		money.Add(int64(-saveAmount))
		bank.Add(int64(saveAmount))
		msg := fmt.Sprintf("Save money: %d \nSave amount in bank: %d\n", money.Load(), bank.Load())
		fmt.Println(msg, money.Load(), bank.Load())
		_, err := w.Write([]byte(msg))
		if err != nil {
			fmt.Println("Error writing response:", err)
		}
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
