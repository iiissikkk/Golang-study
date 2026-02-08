// A map maps keys to values
// The zero value of a map is nil. A nil map has no keys, nor can keys be added
//
// Карты используются для хранения пар ключ:значение. Ключ всегда уникален
// Индексироваться по мапе нельзя!

package main

import "fmt"

func main() {
	weather := map[int]int{
		11: +3,
		12: +2,
		13: +9,
		14: -4,
		15: +1,
	}

	w, ok := weather[30]

	fmt.Println(weather)
	fmt.Println("key[12]: ", weather[12])
	fmt.Println("key[undefined] ", weather[30], ok)
	fmt.Println(w, ok)

	weather[20] = -10
	fmt.Println("new key[20]: ", weather)
	weather[11] = -3
	fmt.Println("change key[11]: ", weather)

	fmt.Println("---------------------------")

	fmt.Println(weather)
	for key, _ := range weather {
		weather[key] += 1
	}
	fmt.Println(weather)

	fmt.Println("---------------------------")

	criminal := map[string]bool{
		"Вася": true,
		"Петя": false,
		"Ваня": true,
		"Вова": true,
	}
	fmt.Println(criminal)

	c, ok := criminal["Вася"]
	if !ok {
		fmt.Println("Человека нет в базе")
		return
	}
	fmt.Println("We got em!")
	if c {
		fmt.Println("Guilty!")
	} else {
		fmt.Println("Not guilty!")
	}
	fmt.Println(c, ok)
}
