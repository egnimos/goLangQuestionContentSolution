/*
QUESTION=>
Using the STRUCT “person”, using a composite literal create a value of type person and assign it to a variable with the
identifier “p1”; print out “p1”; print out just the field fName for “p1”
*/

package main

import (
	"fmt"
)

type person struct {
	FirstName string
	LastName  string
}

func main() {
	p1 := person{
		FirstName: "Sona",
		LastName:  "Tiwa",
	}

	fmt.Println(p1.FirstName)
}
