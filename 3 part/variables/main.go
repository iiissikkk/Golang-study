package main

import (
	"fmt"
	"os"
)

func main() {
	val := os.Getenv("some_var")
	if val != "" {
		fmt.Println("val:", val)
	} else {
		fmt.Println("no val")
	}
}
