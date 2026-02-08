package main

import (
	"fmt"
	"pm/greeting"
	"pm/user"
)

func main() {
	greeting.SayHello()
	greeting.CurseWord()

	u := user.NewUser("Islam", 24)
	fmt.Println(u.GetAge())
}
