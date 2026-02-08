package methods

import (
	"fmt"
	"math/rand"
)

type Bank struct{}

func NewBank() Bank {
	return Bank{}
}

func (c Bank) Pay(usd int) int {
	fmt.Println("Pay with Bank")
	fmt.Printf("Bank USD = %d\n", usd)

	id := rand.Int()
	return id
}

func (c Bank) Cancel(id int) {
	fmt.Println("Cancel Bank with ID:", id)
}
