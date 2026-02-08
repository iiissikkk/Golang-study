package methods

import (
	"fmt"
	"math/rand"
)

type PayPal struct{}

func NewPayPal() PayPal {
	return PayPal{}
}

func (c PayPal) Pay(usd int) int {
	fmt.Println("Pay with PayPal")
	fmt.Printf("PayPal USD = %d\n", usd)

	id := rand.Int()
	return id
}

func (c PayPal) Cancel(id int) {
	fmt.Println("Cancel PayPal with ID:", id)
}
