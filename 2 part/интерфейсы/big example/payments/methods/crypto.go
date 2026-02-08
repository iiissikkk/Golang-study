package methods

import (
	"fmt"
	"math/rand"
)

type Crypto struct{}

func NewCrypto() Crypto {
	return Crypto{}
}

func (c Crypto) Pay(usd int) int {
	fmt.Println("Pay with Crypto")
	fmt.Printf("Crypto USD = %d\n", usd)

	id := rand.Int()
	return id
}

func (c Crypto) Cancel(id int) {
	fmt.Println("Cancel Crypto with ID:", id)
}
