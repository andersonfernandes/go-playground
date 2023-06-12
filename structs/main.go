package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(newName string) {
	(*p).firstName = newName
}

func main() {
	anderson := person{
		firstName: "Anderson",
		lastName:  "Fernandes",
		contact: contactInfo{
			email:   "fernandesanderson14@gmail.com",
			zipCode: 57037620,
		},
	}

	anderson.updateName("Novo Anderson")
	anderson.print()
}
