// An interface type is defined as a set of method signatures
// A value of interface type can hold any value that implements those methods
//
// Интерфейс представляет из себя контракт, определяющий методы, которым обязан обладать объект, удовлетворяющий интерфейсу
// Интерфейсы не содержат конкретной реализации, они только задают обязательные методы, которые должны присутствовать у типа, чтобы он удовлетворял интерфейсу
// Тип реализует интерфейс, предоставляя свою собственную реализацию этих методов

package main

import (
	"fmt"
)

type Auto interface {
	StepOnGas()
	SetpOnBrake()
}

type BMW struct{}

func (b BMW) StepOnGas() {
	fmt.Println("Step on bmw gas")
}
func (b BMW) SetpOnBrake() {
	fmt.Println("Step on bmw brake")
}

type Zhiga struct{}

func (z Zhiga) StepOnGas() {
	fmt.Println("Step on zhiga gas")
}

func rideBMW(bmw BMW) {
	fmt.Println("I am a rider")
	fmt.Println("Step on gas")
	bmw.StepOnGas()
}
func rideZhiga(zhiga Zhiga) {
	fmt.Println("I am a rider")
	fmt.Println("Step on gas")
	zhiga.StepOnGas()
}
func ride(auto Auto) {
	fmt.Println("I am a rider")
	auto.StepOnGas()
	auto.SetpOnBrake()
}

func main() {
	bmw := BMW{}
	bmw.StepOnGas()

	zhiga := Zhiga{}
	zhiga.StepOnGas()

	rideBMW(bmw)
	rideZhiga(zhiga)

	fmt.Println("-------------------")

	ride(bmw)
}
