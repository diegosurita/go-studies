package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (personPointer *person) updateName(newFirstName string) {
	(*personPointer).firstName = newFirstName
}
