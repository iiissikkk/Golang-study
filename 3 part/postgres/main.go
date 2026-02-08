package main

import (
	"fmt"
	"postgresPractice/feature1"
	"postgresPractice/feature2"
	"postgresPractice/simple_connection"
)

func main() {
	fmt.Println("Hello git!")
	feature1.Feature1()
	feature2.Feature2()

	simple_connection.CheckConnection()
}
