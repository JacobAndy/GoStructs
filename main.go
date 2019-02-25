package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName:  "Halpert",
		contactInfo: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 55575,
		},
	}
	// & is an operator, look at this variable and give me the access to the
	// memory address that this variable is pointing at
	// & it is just giving access
	// to the memory storage that jim refers too
	// not the actual value of jim
	jim.updateName("Jimmy")

	jim.print()
}

// * where a type should be, means this can only be called with a pointer
func (pointerToPerson *person) updateName(nFn string) {
	// *pointerToPerson is the actual struct value of jim
	(*pointerToPerson).firstName = nFn
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

//*********** EXTRA NOTES *********** //
// Turn ADDRESS into VALUE with *address
//Turn VALUE into ADDRESS with &value
