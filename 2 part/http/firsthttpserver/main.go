package main

import (
	"fmt"
	"net/http"
	"time"
)

func payHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Payment was added"))
	if err != nil {
		fmt.Println("Something went wrong! Payment wasn't added", err.Error())
	} else {
		fmt.Println("Everything is okay! Payment was added")
	}
}

func cancelHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Payment cancelled"))
	if err != nil {
		fmt.Println("Something went wrong! Payment wasn't cancelled", err.Error())
	} else {
		fmt.Println("Everything is okay! Payment cancelled")
	}
}

func sleepHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)

	_, err := w.Write([]byte("HTTP response!"))
	if err != nil {
		fmt.Println("err:", err)
		return
	}

	fmt.Println("SleepHandler отработал успешно!")
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello World"))
	if err != nil {
		fmt.Println("Something went wrong", err.Error())
	} else {
		fmt.Println("Everything is okay!")
	}
}

func main() {
	http.HandleFunc("/default", handler)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/cancel", cancelHandler)
	http.HandleFunc("/sleep", sleepHandler)

	fmt.Println("Http server is running on port :9091")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Something went wrong!", err.Error())
	}
	fmt.Println("Http server is stopped")
}
