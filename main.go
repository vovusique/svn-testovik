package main

import "fmt"

func main() {
	// just a small comment
	fmt.Printf("Hello, world!!!!!!\n")
	fmt.Printf("I'm starting to make change here\n")

	// that is old code
	fmt.Printf("Now there's a super-duper-muper updated line\n")
}

func newfunc() {
	fmt.Printf("Entierly new thing\n")
	fmt.Printf("Doing my separate stuff\n")
	fmt.Printf("I'm adding more legacy stuff\n")
}

func newfunc2() {
	fmt.Printf("I'm adding shiny new code and not touching legacy\n")
	fmt.Printf("Doing separate stuff after bunch of changes and conflicts\n")
}