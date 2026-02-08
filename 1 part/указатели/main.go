//Go has pointers. A pointer holds the memory address of a value.
//
//The type *T is a pointer to a T value. Its zero value is nil.
//
// Возможность избавиться от необхожимости копирования больших объектов
// в оперативной памяти избегая копирования и заполнения памяти
//
// & - берем адрес у переменной / * - ожидаем указатель на тип данных /  *something - разименовываем указатель

package main

import "fmt"

func main() {
	number := 5
	pointer := &number
	foo(pointer)
	fmt.Println(number)

	s := "Ислам"
	ptr := &s
	boo(ptr)
	fmt.Println(s)

}

func foo(n *int) {
	fmt.Println(n)
	fmt.Println(*n)

	*n = 10
}

func boo(s *string) {
	fmt.Println(s)
	fmt.Println(*s)

	*s = "Zhopa"
}
