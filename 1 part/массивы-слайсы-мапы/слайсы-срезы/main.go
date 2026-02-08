// A slice is a dynamically-sized, flexible view into the elements of an array
//
// len - сколько элементов сейчас в слайсе
// cap - сколько элементов вмещается всего (увеличивается на 2 раза от изначального вмещения)

package main

import "fmt"

type User struct {
	Name    string
	Rating  float64
	Premium bool
}

func main() {
	userArray := []User{
		User{
			Name:    "IslamArray",
			Rating:  4.0,
			Premium: true,
		},
		User{
			Name:    "ZhopaArray",
			Rating:  1.0,
			Premium: true,
		},
		User{
			Name:    "PeterArray",
			Rating:  3.3,
			Premium: false,
		},
	}

	fmt.Println("len: ", len(userArray))
	fmt.Println("cap: ", cap(userArray))
	fmt.Println("userArray: ", userArray)

	userArray = append(
		userArray,
		User{
			Name:    "IslamAppend",
			Rating:  5.0,
			Premium: true,
		})
	fmt.Println("len: ", len(userArray))
	fmt.Println("cap: ", cap(userArray))
	fmt.Println("userArray: ", userArray)

	fmt.Println("---------------------------")

	intSlice := make([]int, 0, 4)
	fmt.Println(intSlice)

	intSlice = append(intSlice, 10)
	intSlice = append(intSlice, 15)
	intSlice = append(intSlice, 25)
	intSlice = append(intSlice, 100)
	fmt.Println(intSlice)

	fmt.Println("len: ", len(intSlice))
	fmt.Println("cap: ", cap(intSlice))
	fmt.Println("intSlice: ", intSlice)
}
