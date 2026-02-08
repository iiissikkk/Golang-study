package main

import (
	"payment/payments"
	"payment/payments/methods"

	"github.com/k0kubun/pp"
)

func main() {
	method := methods.NewBank()

	paymentModule := payments.NewPaymentModule(method)
	paymentModule.Pay("Burger", 5)
	idPhone := paymentModule.Pay("Phone", 500)
	idGame := paymentModule.Pay("game", 150)

	paymentModule.Cancel(idPhone)

	allInfo := paymentModule.AllInfo()
	pp.Println("All your payments: ", allInfo)

	gameInfo := paymentModule.Info(idGame)
	pp.Println("Game info: ", gameInfo)
}
