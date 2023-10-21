package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contactInfo
}

type contactInfo struct {
	email   string
	zipCode int
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Smith",
		contactInfo: contactInfo{
			email:   "jsmith@gmail.com",
			zipCode: 90210},
	}
	// Print out the full name of Jim Smith
	jim.updateFirstname("Jimmy", "Smitthy")
	jim.print()
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateFirstname(newFirstName, newLastname string) {
	p.firstName = newFirstName
	p.lastName = newLastname
}
