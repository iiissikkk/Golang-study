package main

import (
	"docker/http_server"
	"fmt"
)

func main() {
	fmt.Println("Starting server")

	err := http_server.StartHttpServer()
	if err != nil {
		fmt.Println("Something went wrong", err)
	} else {
		fmt.Println("Stop server")
	}
}
