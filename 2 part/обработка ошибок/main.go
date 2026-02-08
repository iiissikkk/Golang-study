package main

import (
	"errors"
	"fmt"
	"math/rand"

	"github.com/k0kubun/pp"
)

type User struct {
	Name     string
	Ballance int
}

func Pay(user *User, usd int) error {
	if user.Ballance-usd < 0 {
		err := errors.New("No money!")
		return err
	}
	user.Ballance -= usd
	return nil
}

type Car struct {
	Armor int
}

func (c *Car) Gas() (int, error) {
	if c.Armor-10 <= 0 {
		return 0, errors.New("Armor is too low")
	}
	kmch := rand.Intn(150)
	c.Armor -= 10

	return kmch, nil
}

func main() {
	user := User{
		Name:     "Islam",
		Ballance: 10,
	}

	pp.Println("User before: ", user)
	err := Pay(&user, 60)
	pp.Println("User after: ", user)

	if err != nil {
		fmt.Println("User not payed!", err.Error())
	} else {
		fmt.Println("User payed!")
	}

	fmt.Println("----------------------------")

	car := Car{
		Armor: 25,
	}

	for {
		pp.Println("Car before: ", car)
		kmch, err := car.Gas()
		if err != nil {
			fmt.Println("Something with speed!", err.Error())
			break
		}
		fmt.Println("Our speed: ", kmch)
	}
}
