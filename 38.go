package main

import "fmt"

type person struct {
	firstName  string
	secondName string
	contactInfo
	dog
}
type dog struct {
	name string
	breed string
}
type contactInfo struct {
	email string
	zip   int
}

func (p person) printPerson() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func main() {
	alex := person{
		firstName:  "Alex",
		secondName: "Anderson",
		contactInfo: contactInfo{
			email: "AlexAnderson@gmail.com",
			zip:   90210,
		},
	}
	alex.updateName("Alox")
	alex.printPerson()
}
