package main

import "fmt"

type person struct {
	firstName  string
	secondName string
	contactInfo
}

type contactInfo struct {
	email string
	zip   int
}

func (p person) printPerson() {
	fmt.Printf("%+v\n", p)
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
	fmt.Println(alex)
	var alexsDog person
	fmt.Println(alexsDog)
	alex.printPerson()
	alexsDog.printPerson()
	alexsDog.firstName = "Fido"
	alexsDog.printPerson()
}
