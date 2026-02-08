package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Payment struct {
	Description string `json:"description"`
	USD         int    `json:"usd"`
	FullName    string `json:"fullName"`
	Address     string `json:"address"`
	Time        time.Time
}

func (p Payment) Println() {
	fmt.Println("Description:", p.Description)
	fmt.Println("USD:", p.USD)
	fmt.Println("FullName:", p.FullName)
	fmt.Println("Address:", p.Address)
	fmt.Println("Time:", p.Time)
}

var (
	mtx            sync.Mutex
	money          = 1000
	paymentHistory = make([]Payment, 0)
)

type HttpResponse struct {
	Money          int       `json:"money"`
	PaymentHistory []Payment `json:"paymentHistory"`
}

func payHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		msg := "Method not allowed. Use POST"
		fmt.Println(msg)
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte(msg))
		return
	}

	var payment Payment
	if err := json.NewDecoder(r.Body).Decode(&payment); err != nil {
		msg := "Bad request: invalid JSON"
		fmt.Println("decode err:", err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte(msg))
		return
	}

	payment.Time = time.Now()
	fmt.Println("OK: correct method POST")
	payment.Println()

	mtx.Lock()
	defer mtx.Unlock()

	if money-payment.USD >= 0 {
		money -= payment.USD
		paymentHistory = append(paymentHistory, payment)
	} else {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("not enough money"))
		return
	}

	httpResponse := HttpResponse{
		Money:          money,
		PaymentHistory: paymentHistory,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	b, err := json.Marshal(httpResponse)
	if err != nil {
		fmt.Println("marshal err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		fmt.Println("write err:", err)
		return
	}
}

func getPaymentHistory(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		msg := "Method not allowed. Use GET"
		fmt.Println(msg)
		w.WriteHeader(http.StatusMethodNotAllowed)
		_, _ = w.Write([]byte(msg))
		return
	}

	mtx.Lock()
	defer mtx.Unlock()

	httpResponse := HttpResponse{
		Money:          money,
		PaymentHistory: paymentHistory,
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	b, err := json.Marshal(httpResponse)
	if err != nil {
		fmt.Println("marshal err:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = w.Write(b)
	if err != nil {
		fmt.Println("write err:", err)
		return
	}
}

func main() {
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/get", getPaymentHistory)

	fmt.Println("Listening on :9091")
	if err := http.ListenAndServe(":9091", nil); err != nil {
		fmt.Println("Ошибка во время работы HTTP сервера:", err)
	}
}
