//Structs
// A struct is a collection of fields.
//
// Несколько логических переменных объеденены в одну сущность

package main

import "fmt"

type User struct {
	// Не должно быть пустым
	Name string // ""
	// Возраст болжен быть больше 0 и меньше 150
	Age int // 0
	// Не должно быть пустым
	PhoneNumber string // ""
	// Закрыт или открыт
	IsClose bool // false
	// Рейтинг должен быть больше либо равен 0 либо меньше либо равен 10
	Rating float64 // 0.0
}

func NewUser(
	name string,
	age int,
	phoneNumber string,
	isClose bool,
	rating float64,
) User {
	if name == "" {
		return User{}
	}
	if age <= 0 || age >= 150 {
		return User{}
	}
	if phoneNumber == "" {
		return User{}
	}
	if rating < 0.0 || rating > 10.0 {
		return User{}
	}

	return User{
		Name:        name,
		Age:         age,
		PhoneNumber: phoneNumber,
		IsClose:     isClose,
		Rating:      rating,
	}
}

func (u *User) CnahgeName(newName string) {
	if newName != "" {
		u.Name = newName
	}
}

func (u *User) CnahgeAge(newAge int) {
	if newAge > 0 && newAge < 150 {
		u.Age = newAge
	}
}

func (u *User) CnahgePhoneNumber(newPhoneNumber string) {
	if newPhoneNumber != "" {
		u.PhoneNumber = newPhoneNumber
	}
}

func (u *User) CloseAccount() {
	u.IsClose = true
}
func (u *User) OpenAccount() {
	u.IsClose = false
}

func (u *User) RatingUp(rating float64) {
	if u.Rating+rating <= 10.0 {
		u.Rating += rating
	}
}
func (u *User) RatingDown(rating float64) {
	if u.Rating-rating >= 10.0 {
		u.Rating -= rating
	}
}

//func (user User) Greeting() {
//	fmt.Println("Hello!")
//	fmt.Printf("My name is %s!\n", user.Name)
//	fmt.Println("")
//}

//func (user User) Goodbye() { // ресивер копий
//	fmt.Println("Bye!")
//	fmt.Printf("My name was %s!\n", user.Name)
//}

//func RatingUp(user *User, rating float64) {
//	if user.Rating+rating <= 10.0 {
//		user.Rating += rating
//		fmt.Printf("Rating is %f, you're up!", user.Rating)
//		fmt.Println("\n")
//	} else {
//		fmt.Printf("Rating is %f, nothing changed :(", user.Rating)
//		fmt.Println("\n")
//	}
//}

//func (user *User) RatingUp(rating float64) { // (user *User) ресивер по указателю
//	if user.Rating+rating <= 10.0 {
//		user.Rating += rating
//		fmt.Printf("Rating is %f, you're up!", user.Rating)
//		fmt.Println("\n")
//	} else {
//		fmt.Printf("Rating is %f, nothing changed :(", user.Rating)
//		fmt.Println("\n")
//	}
//}

func main() {
	//user := User{
	//	Name:        "Islam",
	//	Age:         24,
	//	PhoneNumber: "+7(928)711-00-42",
	//	IsClose:     true,
	//	Rating:      4.7,
	//}
	//
	//// ptr := &user
	////fmt.Println(user, "\n")
	////user.Greeting()
	////user.RatingUp(1.5)
	//// RatingUp(ptr, 2.0)
	////user.Goodbye()
	//
	//fmt.Println("Before: ", user.Rating)
	////x := 0.0
	////fmt.Scan(&x)
	////user.RatingUp(x)
	//user.RatingUp(5.5)
	//fmt.Println("After: ", user.Rating)

	user := NewUser(
		"Islam",
		24,
		"+7(928)711-00-42",
		false,
		4.5,
	)
	fmt.Println(user)
}
