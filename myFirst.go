package main

import "fmt"

func main() {
	fmt.Println("Welcome to Go-Lang!")
	fmt.Print("Enter your name to introduce: ")
	var name string
	fmt.Scan(&name)
	name = name + "!"
	fmt.Println("Hi,", name)
	fmt.Println("Nice to meet you. I'm Go ")
}
