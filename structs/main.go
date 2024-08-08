package main

import "fmt"

func main() {

	type User struct {
		name, address string
		age           int
	}

	today := User{
		name:    "magic",
		address: "seoul",
		age:     25,
	}

	fmt.Println("today data: ", today.name, today.address, today.age)
	today.age = 30
	fmt.Println("today data: ", today.name, today.address, today.age)

}
