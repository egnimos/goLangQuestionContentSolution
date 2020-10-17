/*
QUESTION=>
Create a new type called person which is STRUCT that stores fName and lName
*/

package main

import (
	"fmt"
)

func main() {

	person := struct {
		FirstName string
		LastName  string
	}{
		"Niteesh",
		"Dubey",
	}

	fmt.Println(person)
}
