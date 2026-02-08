package greeting

import "fmt"

func SayHello() {
	fmt.Println("Hello golang packages! How many?")
	x := GiveMeIn()
	fmt.Println(x)
}
