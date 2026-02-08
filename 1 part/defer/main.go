// A defer statement defers the execution of a function until the surrounding function returns.
//
//The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
//
//

package main

import "fmt"

func main() {
	fmt.Println("Hello - func main")

	defer func() {
		fmt.Println("func main - Выполнюсь в конце")
	}()

	hello()
	multiDefer()
}

func hello() {
	fmt.Println("Hello - func hello")
	defer func() {
		fmt.Println("func hello - Выполнюсь в конце")
	}()
	fmt.Println("Hello func end")
}

func multiDefer() {
	defer func() {
		fmt.Println("Defer-1")
	}()
	defer func() {
		fmt.Println("Defer-2")
	}()
	defer func() {
		fmt.Println("Defer-3")
	}()
	fmt.Println("func multiDefer")
}
