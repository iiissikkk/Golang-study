// Массив последовательность значений которые находятся в памяти непрерывно
// An array's length is part of its type, so arrays-slices-map cannot be resized
//
// Нумерация в массиве начинается с 0

package main

import (
	"fmt"

	"github.com/k0kubun/pp"
)

type User struct {
	Name    string
	Rating  float64
	Premium bool
}

func main() {
	arr := [5]int{5, 66, 7, 100, 1}
	fmt.Println("arr: ", arr)
	fmt.Println("arr[2]: ", arr[2])
	arr[2] += 5
	fmt.Println("arr[2]+5: ", arr[2])
	fmt.Println("arr: ", arr)

	arr2 := [6]int{}
	arr2[0] = arr[0]
	arr2[1] = arr[1]
	arr2[2] = arr[2]
	arr2[3] = arr[3]
	arr2[4] = arr[4]
	arr2[5] = 77
	fmt.Println("arr2: ", arr2)

	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 0 {
			arr[i] *= 2
		}
		fmt.Println(i, " - ", arr[i])
	}
	fmt.Println("arr[i]%2==0 * 2: ", arr)

	fmt.Println("---------------------------")

	user1 := User{
		Name:    "Islam",
		Rating:  4.0,
		Premium: true,
	}
	user2 := User{
		Name:    "Zhopa",
		Rating:  1.0,
		Premium: true,
	}
	user3 := User{
		Name:    "Peter",
		Rating:  3.3,
		Premium: false,
	}

	if user1.Premium {
		user1.Rating += 1.0
	}
	if user2.Premium {
		user2.Rating += 1.0
	}
	if user3.Premium {
		user3.Rating += 1.0
	}

	pp.Println(user1)
	pp.Println(user2)
	pp.Println(user3)

	fmt.Println("---------------------------")

	userArray := [3]User{
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

	for i := 0; i < len(userArray); i++ {
		if userArray[i].Premium {
			userArray[i].Rating += 1.0
		}
	}
	pp.Println(userArray)

	fmt.Println("---------------------------")

	// вывод на экран вторым видом for
	for index, value := range userArray {
		fmt.Println(index, "->", value)
	}
	fmt.Println("---------------------")

	// просто вывод всех индексов
	for index, _ := range userArray {
		fmt.Println("index:", index)
	}
	fmt.Println("---------------------")

	// просто вывод всех значений
	for _, value := range userArray {
		fmt.Println("value:", value)
	}
	fmt.Println("---------------------")

	// прибавляет каждому элементу, который больше 60, единицу вторым видом for
	for index, _ := range userArray {
		if userArray[index].Premium {
		}
		userArray[index].Rating += 1

		// value++ не сработало бы
		// так как value это всего-лишь копия значения из массива
	}
	pp.Println(userArray)
}
