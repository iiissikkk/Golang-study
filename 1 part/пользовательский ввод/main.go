package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Do something: ")
		ok := scanner.Scan()
		if !ok {
			fmt.Println("Error!")
			return
		}
		text := scanner.Text()
		fields := strings.Fields(text)

		if len(fields) == 0 {
			fmt.Println("Error! Nothing to do!")
			return
		}

		cmd := fields[0]

		if cmd == "Exit" {
			fmt.Println("Bye!")
			return
		}

		if cmd == "Add" {
			str := ""
			for i := 1; i < len(fields); i++ {
				str += fields[i]
				if i != len(fields)-1 {
					str += " "
				}
			}
			fmt.Println("Do you wanna add: ", str)
			fmt.Println("")
		} else if cmd == "Del" {
			str := ""
			for i := 1; i < len(fields); i++ {
				str += fields[i]
				if i != len(fields)-1 {
					str += " "
				}
			}
			fmt.Println("Do you wanna delete: ", str)
			fmt.Println("")
		} else if cmd == "Help" {
			fmt.Println("Command: Add {what we need to add}")
			fmt.Println("-- this command may help you add something --")
			fmt.Println("")
			fmt.Println("Command: Del {what we need to delete}")
			fmt.Println("-- this command may help you add something --")
			fmt.Println("")
			fmt.Println("Command: Help ")
			fmt.Println("-- some text --")
			fmt.Println("")
		} else {
			fmt.Println("User invalid!")
		}
	}
}
