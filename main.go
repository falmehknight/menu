package main

import "fmt"

func main() {
	var cmd string
	for {
		fmt.Print("Please input a command: ")
		fmt.Scanln(&cmd)
		switch cmd {
		case "h":
			fmt.Println("This is help command.")
		case "q":
			fmt.Println("bye.")
			return
		default:
			fmt.Println("Wrong command!")
		}
	}
}
