/*
QUESTION=>
Initialize a MAP using a composite literal where the key is a string and the
value is an int; print out the map; range over the map printing out just the
key; range over the map printing out both the key and the value
*/

package main

import (
	"fmt"
)

func main() {
	mp := map[string]int{
		"Niteesh": 1604610055,
		"Rahul":   1604610063,
		"Pragya":  1604610056,
		"Nerraj":  1604610054,
		"Monu":    1604610052,
	}

	//print the keys
	for i, _ := range mp {
		fmt.Println("So the keys are: ", i)
	}

	//print the keys and values
	for i, v := range mp {
		fmt.Println("So the index is: ", i, "--- And the values are: ", v)
	}
}
