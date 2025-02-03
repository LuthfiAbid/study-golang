package main

import "fmt"

type contactInfo struct {
	address string
	zipCode int
}

type person struct {
	fName string
	lName string
	contactInfo
}

func main() {
	fPerson := person{
		fName: "Luthfi",
		lName: "Abid",
		contactInfo: contactInfo{
			address: "Kota Batu",
			zipCode: 65316,
		},
	}

	fPerson.updateName("Muhammad")
	fPerson.print()

}

func (pointerToPerson *person) updateName(fName string) {
	(*pointerToPerson).fName = fName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
